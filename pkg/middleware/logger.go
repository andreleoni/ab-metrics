package middleware

import (
	"encoding/hex"
	"log/slog"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func DefaultStructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logCorrelationID, _ := randomHex(10)
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Set("logCorrelationID", logCorrelationID)

		c.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
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

		if c.Writer.Status() >= 500 {
			slog.Error(
				param.ErrorMessage,
				"client_id", param.ClientIP,
				"method", param.Method,
				"status_code", param.StatusCode,
				"body_size", param.BodySize,
				"path", param.Path,
				"latency", param.Latency.String(),
				"correlation_id", logCorrelationID,
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
				"correlation_id", logCorrelationID,
			)
		}
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
