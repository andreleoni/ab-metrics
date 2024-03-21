package persistence

import (
	"ab-metrics/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type ExperimentRepository struct {
	db *mongo.Client
}

func NewExperimentRepository(db *mongo.Client) ExperimentRepository {
	return ExperimentRepository{db: db}
}

func (ExperimentRepository) GetByKey(key string) entity.Experiment {
	variationA := entity.Variation{ID: "foovariation", ExperimentID: "apskd", Key: "a", Percentage: 33}
	variationB := entity.Variation{ID: "barvariation", ExperimentID: "apskd", Key: "b", Percentage: 33}
	variationC := entity.Variation{ID: "foovariation", ExperimentID: "apskd", Key: "control", Percentage: 34}

	experiment := entity.Experiment{
		ID:         "fooexperiment",
		Name:       "first experiment",
		Key:        "first_experiment",
		Status:     "active",
		Variations: []entity.Variation{variationA, variationB, variationC}}

	return experiment
}
