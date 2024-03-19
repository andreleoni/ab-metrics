package main

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"ab-metrics/internal/domain/service"
	"ab-metrics/internal/infrastructure/database/sqlite"
	"ab-metrics/internal/infrastructure/persistence"
	"ab-metrics/internal/usecase"
	"ab-metrics/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	desiredLogLevel := os.Getenv("LOG_LEVEL")

	logLevel := slog.LevelInfo

	if desiredLogLevel == "DEBUG" {
		logLevel = slog.LevelDebug
	} else if desiredLogLevel == "WARN" {
		logLevel = slog.LevelWarn
	} else if desiredLogLevel == "ERROR" {
		logLevel = slog.LevelError
	}

	opts := &slog.HandlerOptions{Level: logLevel}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	sqlite.SQLiteSetup()

	r := gin.New() // empty engine

	r.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	r.Use(gin.Recovery())                       // adds the default recovery middleware

	experimentRepositoryImpl := persistence.NewExperimentRepository(sqlite.Sqlite)
	actorRepositoryImpl := persistence.NewActorRepository(sqlite.Sqlite)
	goalRepositoryImpl := persistence.NewGoalRepository(sqlite.Sqlite)

	scenarioEligibilityServiceImpl := service.NewScenarioEligibilityService()

	v1 := r.Group("/api/v1")
	// POST /api/v1/experiment/:scenario
	//   request: { "device_uuid": "test" }
	//   response: { "id": "test", "scenario": "" }
	//
	//   gerar um token baseado no uuid, pegar a porcentagem ativa e escolher o cenário para o usuário
	v1.POST("/experiment/:experiment", func(c *gin.Context) {
		logCorrelationID, logCorrelationIDExists := c.Get("logCorrelationID")

		experiment, experimentExists := c.Params.Get("experiment")
		if !experimentExists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "experiment not found",
			})

			return
		}

		// Parse body
		bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
		jsonMap := make(map[string]string)
		json.Unmarshal(bodyAsByteArray, &jsonMap)

		identifier := jsonMap["identifier"]
		if identifier == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "identifier not found",
			})

			return
		}

		contextlogger := logger

		if logCorrelationIDExists {
			contextlogger = logger.With("correlation_id", logCorrelationID)
		}

		contextlogger = logger.With("identifier", identifier)

		geiuci := usecase.ScenarioEligibilityCheckInput{
			Identifier: identifier,
			Experiment: experiment}

		getExperimentUserCase := usecase.NewScenarioEligibilityCheckUseCase(
			contextlogger,
			experimentRepositoryImpl,
			actorRepositoryImpl,
			scenarioEligibilityServiceImpl)

		output, err := getExperimentUserCase.Execute(geiuci)

		if err != nil {
			contextlogger.Error(err.Error(),
				"identifier", identifier,
				"experiment", experiment,
			)

			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, output)
	})

	// POST /api/v1/experiment/:goal
	//   request: { "id": "tzest", "goal": "checkout_finished" }
	//   response: { "id": "test", "scenario": "" }
	//
	//   pegar o experiment ID, e cadastrar o goal para ele caso estiver presente
	v1.POST("/experiment/goal/:goal", func(c *gin.Context) {
		logCorrelationID, logCorrelationIDExists := c.Get("logCorrelationID")

		contextlogger := logger

		if logCorrelationIDExists {
			contextlogger = logger.With("correlation_id", logCorrelationID)
		}

		goal, goalExists := c.Params.Get("goal")
		if !goalExists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "goal not found",
			})

			return
		}

		// Parse body
		bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
		jsonMap := make(map[string]string)
		json.Unmarshal(bodyAsByteArray, &jsonMap)

		identifier := jsonMap["identifier"]
		if identifier == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "identifier not found",
			})

			return
		}

		createActorGoalCheckUseCase := usecase.NewCreateActorGoalCheckUseCase(
			contextlogger, actorRepositoryImpl, goalRepositoryImpl,
		)

		recordGoalInput := usecase.CreateActorGoalCheckInput{ActorID: identifier, GoalKey: goal}

		goalOutput, err := createActorGoalCheckUseCase.Execute(recordGoalInput)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": goalOutput.Status,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
