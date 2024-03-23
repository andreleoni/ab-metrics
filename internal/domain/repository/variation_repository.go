package repository

import (
	"ab-metrics/internal/domain/entity"
)

type VariationRepository interface {
	GetByExperimentKey(key string) ([]entity.Variation, error)
	Create(v *entity.Variation) error
}
