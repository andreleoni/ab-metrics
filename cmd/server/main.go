package main

import (
	"log/slog"
	"net/http"
	"os"

	"ab-metrics/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// POST /api/v1/experiment/:scenario
	//   request: { "device_uuid": "test" }
	//   response: { "id": "test", "scenario": "" }
	//
	//   gerar um token baseado no uuid, pegar a porcentagem ativa e escolher o cenário para o usuário

	// POST /api/v1/experiment/:goal
	//   request: { "id": "tzest", "goal": "checkout_finished" }
	//   response: { "id": "test", "scenario": "" }
	//
	//   pegar o experiment ID, e cadastrar o goal para ele caso estiver presente

	// GET /api/v1/stats/:scenario
	//   response: { "groups": { "a": 1, "b": 2 } }

	r := gin.New() // empty engine

	r.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	r.Use(gin.Recovery())                       // adds the default recovery middleware

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
