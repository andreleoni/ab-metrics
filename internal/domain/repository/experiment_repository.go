package repository

import (
	"ab-metrics/internal/domain/entity"
)

type ExperimentRepository interface {
	GetByKey(key string) (entity.Experiment, bool, error)
	Create(e *entity.Experiment) error
}
