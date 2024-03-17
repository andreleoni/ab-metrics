package persistence

import (
	"ab-metrics/internal/domain/entity"
	"log/slog"
)

type ActorRepository struct {
}

func NewActorRepository() ActorRepository {
	return ActorRepository{}
}

func (ActorRepository) Create(a entity.Actor) entity.Actor {
	// TODO: create actor

	slog.Info("User created!", "actor", a)

	return a
}
