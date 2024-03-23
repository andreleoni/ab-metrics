package persistence

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/pkg/random"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperimentRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func NewExperimentRepository(db *mongo.Client) ExperimentRepository {
	collection := db.Database("abmetrics").Collection("experiments")

	return ExperimentRepository{db: db, collection: collection}
}

func (er ExperimentRepository) GetByKey(key string) (entity.Experiment, bool, error) {
	experiment := entity.Experiment{}

	filter := bson.M{
		"key": key,
	}

	err := er.collection.FindOne(context.Background(), filter).Decode(&experiment)

	if err != nil && err.Error() == "mongo: no documents in result" {
		return experiment, false, nil
	} else if err != nil {
		slog.Error("Error on retrieve mongodb result", "error", err)
	}

	slog.Debug("ExperimentRepository#GetByKey",
		"error", err,
		"retrieve_goal", experiment)

	return experiment, true, err
}

func (vr ExperimentRepository) Create(e *entity.Experiment) error {
	e.ID = random.Hex(10)
	e.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// Insert the item into the collection
	insertResult, err := vr.collection.InsertOne(context.Background(), e)
	if err != nil {
		slog.Error("Error on insert mongodb result", "error", err)

		return err
	}

	slog.Debug("ExperimentRepository#Create",
		"insertResult", insertResult,
		"error", err,
	)

	return nil
}
