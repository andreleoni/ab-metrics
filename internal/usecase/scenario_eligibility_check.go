package usecase

import (
	"ab-metrics/internal/repository"
	"log/slog"
)

type ScenarioEligibilityCheckUseCase struct {
	logger               *slog.Logger
	experimentRepository repository.Experiment
}

type ScenarioEligibilityCheckInput struct {
	Identifier string
	Scenario   string
}

type ScenarioEligibilityCheckOuput struct {
	Token     string `json:"token"`
	Variation string `json:"variation"`
}

func NewScenarioEligibilityCheckUseCase(logger *slog.Logger) ScenarioEligibilityCheckUseCase {
	return ScenarioEligibilityCheckUseCase{logger: logger}
}

func (geiuc ScenarioEligibilityCheckUseCase) Execute(
	geiuci ScenarioEligibilityCheckInput) (ScenarioEligibilityCheckOuput, error) {

	experiments := geiuc.experimentRepository.Get()

	geiuc.logger.Debug("Experiments result", "experiments", experiments)

	return ScenarioEligibilityCheckOuput{Token: "pokasd", Variation: "paoskd"}, nil
}
