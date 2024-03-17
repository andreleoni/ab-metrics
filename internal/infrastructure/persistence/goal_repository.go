package persistence

import (
	"ab-metrics/internal/domain/entity"
	"fmt"

	"ab-metrics/pkg/random"
)

type GoalRepository struct {
}

func NewGoalRepository() GoalRepository {
	return GoalRepository{}
}

func (GoalRepository) Get(actorID string, key string) (entity.Goal, error) {
	// TODO: create Goal

	if actorID == "simulatenotfound" {
		return entity.Goal{}, fmt.Errorf("not found")
	}

	return entity.Goal{}, nil
}

func (GoalRepository) Create(g entity.Goal) (entity.Goal, error) {
	// TODO: create Goal

	g.ID = random.Hex(10)

	return g, nil
}
