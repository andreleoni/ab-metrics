package usecase

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/internal/domain/repository"
	"ab-metrics/internal/domain/service"
	"ab-metrics/pkg/random"
	"log/slog"
)

type ScenarioEligibilityCheckUseCase struct {
	logger                     *slog.Logger
	experimentRepository       repository.ExperimentRepository
	actorRepository            repository.ActorRepository
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
	experiment := geiuc.experimentRepository.GetByKey(geiuci.Experiment)

	geiuc.logger.Debug(
		"ScenarioEligibilityCheckOuput#Execute",
		"retrieve_experiment", experiment,
		"input", geiuci)

	seci := service.ScenarioEligibilityServiceInput{Identifier: geiuci.Identifier, Experiment: experiment}

	sesVariation := geiuc.scenarioEligibilityService.GetVariation(seci)

	if sesVariation.Variation.Key == "" {
		geiuc.logger.Debug(
			"ScenarioEligibilityCheckOuput#Execute: variation not found to user",
			"retrieve_experiments", experiment,
			"input", geiuci)

		return ScenarioEligibilityCheckOuput{Variation: "control"}, nil
	}

	variationID := sesVariation.Variation.ID
	actorID := random.Hex(10)

	if sesVariation.Variation.Key != "control" {
		actor := entity.Actor{
			ID:          actorID,
			VariationID: variationID,
			Identifier:  geiuci.Identifier}

		geiuc.actorRepository.Create(actor)
	}

	return ScenarioEligibilityCheckOuput{Token: actorID, Variation: sesVariation.Variation.Key}, nil
}
