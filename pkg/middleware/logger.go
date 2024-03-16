package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func DefaultStructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Log using the params
		if c.Writer.Status() >= 500 {
			slog.Error(
				param.ErrorMessage,
				"client_id", param.ClientIP,
				"method", param.Method,
				"status_code", param.StatusCode,
				"body_size", param.BodySize,
				"path", param.Path,
				"latency", param.Latency.String(),
			)
		} else {
			slog.Info(
				param.ErrorMessage,
				"client_id", param.ClientIP,
				"method", param.Method,
				"status_code", param.StatusCode,
				"body_size", param.BodySize,
				"path", param.Path,
				"latency", param.Latency.String(),
			)
		}
	}
}
