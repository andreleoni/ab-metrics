package repository

import "ab-metrics/internal/domain/entity"

type ActorRepository interface {
	Create(a entity.Actor) (entity.Actor, error)
}
