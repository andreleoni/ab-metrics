package usecase

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/internal/domain/repository"
	"ab-metrics/internal/domain/service"
	"log/slog"
)

type ScenarioEligibilityCheckUseCase struct {
	logger                     *slog.Logger
	experimentRepository       repository.ExperimentRepository
	actorRepository            repository.ActorRepository
	variationRepository        service.VariationRepository
	scenarioEligibilityService service.ScenarioEligibilityService
}

type ScenarioEligibilityCheckInput struct {
	Identifier string
	Experiment string
}

type ScenarioEligibilityCheckOuput struct {
	Token     string `json:"token,omitempty"`
	Variation string `json:"variation"`
}

func NewScenarioEligibilityCheckUseCase(
	logger *slog.Logger,
	experimentRepository repository.ExperimentRepository,
	actorRepository repository.ActorRepository,
	scenarioEligibilityService service.ScenarioEligibilityService) ScenarioEligibilityCheckUseCase {

	return ScenarioEligibilityCheckUseCase{
		logger:                     logger,
		experimentRepository:       experimentRepository,
		actorRepository:            actorRepository,
		scenarioEligibilityService: scenarioEligibilityService}
}

func (geiuc ScenarioEligibilityCheckUseCase) Execute(
	geiuci ScenarioEligibilityCheckInput) (ScenarioEligibilityCheckOuput, error) {
	var err error

	experiment, exists, err := geiuc.experimentRepository.GetByKey(geiuci.Experiment)
	if !exists {
		return ScenarioEligibilityCheckOuput{Variation: "control"}, nil
	}

	variations, err := geiuc.variationRepository.GetByExperimentKey(geiuci.Experiment)
	if !exists {
		return ScenarioEligibilityCheckOuput{Variation: "control"}, nil
	}

	if err != nil {
		geiuc.logger.Error(
			"ScenarioEligibilityCheckOuput#Execute",
			"retrieve_experiment", experiment,
			"input", geiuci,
			"error", err)
	}

	geiuc.logger.Debug(
		"ScenarioEligibilityCheckOuput#Execute",
		"retrieve_experiment", experiment,
		"input", geiuci)

	seci := service.ScenarioEligibilityServiceInput{
		Identifier: geiuci.Identifier,
		UniqueKey:  experiment.Key,
		Variations: variations,
	}

	sesVariation := geiuc.scenarioEligibilityService.GetVariation(seci)

	if sesVariation.Variation.Key == "" {
		geiuc.logger.Debug(
			"ScenarioEligibilityCheckOuput#Execute: variation not found to user",
			"retrieve_experiments", experiment,
			"input", geiuci)

		return ScenarioEligibilityCheckOuput{}, nil
	}

	variationID := sesVariation.Variation.ID

	actor := entity.Actor{}
	actor.VariationID = variationID
	actor.Identifier = geiuci.Identifier

	err = geiuc.actorRepository.Create(&actor)
	if err != nil {
		geiuc.logger.Debug(
			"ScenarioEligibilityCheckOuput#Execute: error on create actor",
			"retrieve_experiments", experiment,
			"input", geiuci,
			"error", err,
		)
	}

	return ScenarioEligibilityCheckOuput{Token: actor.ID, Variation: sesVariation.Variation.Key}, nil
}

// LookupVariationByID searches for a variation in the experiment's variations by ID.
// It returns the found variation and a boolean indicating whether the variation was found.
// func (e ScenarioEligibilityCheckUseCase) LookupVariationByID(vid string) (entity.Variation, bool) {
// 	if v.ID == vid {
// 		return v, true
// 	}

// 	return entity.Variation{}, false
// }
