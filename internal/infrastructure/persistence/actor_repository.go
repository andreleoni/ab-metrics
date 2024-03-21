package persistence

import (
	"ab-metrics/internal/domain/entity"
	"context"
	"log/slog"
	"time"

	"ab-metrics/pkg/random"

	"go.mongodb.org/mongo-driver/mongo"
)

type ActorRepository struct {
	db *mongo.Client
}

func NewActorRepository(db *mongo.Client) ActorRepository {
	return ActorRepository{db: db}
}

func (ar ActorRepository) Create(a *entity.Actor) error {
	a.ID = random.Hex(10)
	a.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// Access the collection
	collection := ar.db.Database("abmetrics").Collection("actors")

	// Insert the item into the collection
	insertResult, err := collection.InsertOne(context.Background(), a)
	if err != nil {
		slog.Error("Error on retrieve mongodb result", "error", err)
	}

	slog.Debug("ActorRepositoryPersistence#Create",
		"insertResult", insertResult,
		"error", err,
	)

	return err
}
