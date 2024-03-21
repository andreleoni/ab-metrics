package persistence

import (
	"ab-metrics/internal/domain/entity"
	"context"
	"log/slog"
	"time"

	"ab-metrics/pkg/random"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoalRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func NewGoalRepository(db *mongo.Client) GoalRepository {
	collection := db.Database("abmetrics").Collection("goals")

	return GoalRepository{db: db, collection: collection}
}

func (gr GoalRepository) Get(actorID string, key string) (entity.Goal, bool, error) {
	goal := entity.Goal{}

	collection := gr.db.Database("abmetrics").Collection("goals")

	filter := bson.M{
		"actorid": actorID,
		"key":     key,
	}

	err := collection.FindOne(context.Background(), filter).Decode(&goal)

	if err != nil && err.Error() == "mongo: no documents in result" {
		return goal, false, nil
	} else if err != nil {
		slog.Error("Error on retrieve mongodb result", "error", err)
	}

	slog.Debug("GoalRepositoryPersistence#Create",
		"error", err,
		"retrieve_goal", goal)

	return goal, true, err
}

func (gr GoalRepository) Create(g *entity.Goal) error {
	g.ID = random.Hex(10)
	g.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// Insert the item into the collection
	insertResult, err := gr.collection.InsertOne(context.Background(), g)
	if err != nil {
		slog.Error("Error on insert mongodb result", "error", err)
	}

	slog.Debug("GoalRepositoryPersistence#Create",
		"insertResult", insertResult,
		"error", err,
	)

	return err
}
