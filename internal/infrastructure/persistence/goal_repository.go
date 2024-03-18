package persistence

import (
	"ab-metrics/internal/domain/entity"

	"ab-metrics/pkg/random"

	"gorm.io/gorm"
)

type GoalRepository struct {
	sqlite *gorm.DB
}

func NewGoalRepository(sqlite *gorm.DB) GoalRepository {
	return GoalRepository{sqlite: sqlite}
}

func (gr GoalRepository) Get(actorID string, key string) (entity.Goal, error) {
	goal := entity.Goal{ActorID: actorID, Key: key}

	result := gr.sqlite.First(&goal)

	return goal, result.Error
}

func (gr GoalRepository) Create(g entity.Goal) error {
	g.ID = random.Hex(10)

	result := gr.sqlite.Create(&g)

	return result.Error
}
