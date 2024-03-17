package persistence

import (
	"ab-metrics/internal/domain/entity"
)

type variation struct {
	percentage int
	key        string
}

var experiments = map[string][]variation{
	"new_checkout_button": {
		{80, "control"},
		{20, "a"},
		{20, "b"},
	},
}

type ExperimentRepository struct{}

func NewExperimentRepository() ExperimentRepository {
	return ExperimentRepository{}
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
