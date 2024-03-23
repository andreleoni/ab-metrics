package persistence

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/pkg/random"
	"context"
	"log"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VariationRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func NewVariationRepository(db *mongo.Client) VariationRepository {
	collection := db.Database("abmetrics").Collection("variations")

	return VariationRepository{db: db, collection: collection}
}

func (vr VariationRepository) GetByExperimentKey(key string) ([]entity.Variation, error) {
	variations := []entity.Variation{}

	filter := bson.M{
		"key": key,
	}

	cursor, err := vr.collection.Find(context.Background(), filter)

	if err := cursor.Err(); err != nil {
		slog.Error("VariationRepository#GetByExperimentKey: error on get cursor", "error", err)
	}

	for cursor.Next(context.Background()) {
		var variation entity.Variation

		if err := cursor.Decode(&variation); err != nil {
			slog.Error("VariationRepository#GetByExperimentKey: error on iterator", "error", err)
		}

		variations = append(variations, variation)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return variations, err
}

func (vr VariationRepository) Create(v *entity.Variation) (*entity.Variation, error) {
	v.ID = random.Hex(10)
	v.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	// Insert the item into the collection
	insertResult, err := vr.collection.InsertOne(context.Background(), v)
	if err != nil {
		slog.Error("Error on insert mongodb result", "error", err)

		return v, err
	}

	slog.Debug("VariationRepository#Create",
		"insertResult", insertResult,
		"error", err,
	)

	return v, nil
}
