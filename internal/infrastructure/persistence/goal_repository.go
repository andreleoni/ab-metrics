package persistence

import (
	"ab-metrics/internal/domain/entity"
	"log/slog"
	"time"

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
	goal := entity.Goal{}

	result := gr.sqlite.Where("actor_id = ? AND key = ?", actorID, key).First(&goal)

	slog.Debug("Goal Get Result", "result", result)

	return goal, result.Error
}

func (gr GoalRepository) Create(g entity.Goal) error {
	g.ID = random.Hex(10)
	g.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	result := gr.sqlite.Create(&g)

	slog.Debug("Goal Create Result",
		"goal", g,
		"result", result)

	return result.Error
}
