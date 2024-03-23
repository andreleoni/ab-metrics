package middleware

import (
	"ab-metrics/internal/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})

			c.Abort()

			return
		}

		keyValidator := service.NewApiKeyValidatorService()

		if !keyValidator.Valid(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})

			c.Abort()

			return
		}

		c.Next()
	}
}
