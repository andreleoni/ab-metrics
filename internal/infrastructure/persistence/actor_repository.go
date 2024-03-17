package persistence

import (
	"ab-metrics/internal/domain/entity"

	"ab-metrics/pkg/random"
)

type ActorRepository struct {
}

func NewActorRepository() ActorRepository {
	return ActorRepository{}
}

func (ActorRepository) Create(a entity.Actor) (entity.Actor, error) {
	// TODO: create actor

	a.ID = random.Hex(10)

	return a, nil
}
