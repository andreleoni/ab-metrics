package main

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"ab-metrics/internal/usecase"
	"ab-metrics/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: add log level here
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	r := gin.New() // empty engine

	r.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	r.Use(gin.Recovery())                       // adds the default recovery middleware

	v1 := r.Group("/api/v1")
	// POST /api/v1/experiment/:scenario
	//   request: { "device_uuid": "test" }
	//   response: { "id": "test", "scenario": "" }
	//
	//   gerar um token baseado no uuid, pegar a porcentagem ativa e escolher o cenário para o usuário
	v1.POST("/experiment/:scenario", func(c *gin.Context) {
		logCorrelationID, logCorrelationIDExists := c.Get("logCorrelationID")

		scenario, scenarioExists := c.Params.Get("scenario")
		if !scenarioExists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "scenario not found",
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

		geiuci := usecase.ScenarioEligibilityCheckInput{Identifier: identifier, Scenario: scenario}

		getExperimentUserCase := usecase.NewScenarioEligibilityCheckUseCase(contextlogger)
		output, err := getExperimentUserCase.Execute(geiuci)

		if err != nil {
			contextlogger.Error(err.Error(),
				"identifier", identifier,
				"scenario", scenario,
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
		contextlogger := logger

		contextlogger.Info("OIEEEEE")

		c.JSON(http.StatusOK, gin.H{
			"message": "qwopkeqpwe",
		})
	})

	// GET /api/v1/stats/:scenario
	//   response: { "groups": { "a": 1, "b": 2 } }
	v1.POST("/stats/:scenario", func(c *gin.Context) {
		contextlogger := logger

		contextlogger.Info("OIEEEEE")

		c.JSON(http.StatusOK, gin.H{
			"message": "kkkk",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
