package service

import (
	"log/slog"
	"os"
)

type ApiKeyValidatorService struct {
	key string
}

func NewApiKeyValidatorService() ApiKeyValidatorService {
	adminApiKey := os.Getenv("ADMIN_API_KEY")

	return ApiKeyValidatorService{key: adminApiKey}
}

func (akvs ApiKeyValidatorService) Valid(key string) bool {
	slog.Debug("ApiKeyValidatorService#Valid",
		"key", key,
		"adminApiKey", akvs,
	)

	return akvs.key == key
}

func (u ApiKeyValidatorService) LogValue() slog.Value {
	return slog.GroupValue(slog.String("key", "[REDACTED]"))
}
