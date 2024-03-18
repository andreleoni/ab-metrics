package persistence

import (
	"ab-metrics/internal/domain/entity"

	"ab-metrics/pkg/random"

	"gorm.io/gorm"
)

type ActorRepository struct {
	sqlite *gorm.DB
}

func NewActorRepository(sqlite *gorm.DB) ActorRepository {
	return ActorRepository{sqlite: sqlite}
}

func (ar ActorRepository) Create(a *entity.Actor) error {
	a.ID = random.Hex(10)

	result := ar.sqlite.Create(&a)

	return result.Error
}
