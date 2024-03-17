package repository

import (
	"ab-metrics/internal/domain/entity"
)

type GoalRepository interface {
	Get(actorID string, key string) (entity.Goal, error)
	Create(entity.Goal) (entity.Goal, error)
}
