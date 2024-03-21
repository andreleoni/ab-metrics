package repository

import (
	"ab-metrics/internal/domain/entity"
)

type GoalRepository interface {
	Get(actorID string, key string) (goal entity.Goal, exists bool, err error)
	Create(*entity.Goal) error
}
