package usecase

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/internal/domain/repository"
	"log/slog"
)

type CreateNewExperimentUseCase struct {
	logger               *slog.Logger
	experimentRepository repository.ExperimentRepository
}

type CreateNewExperimentInput struct {
	Name string
	Key  string
}

type CreateNewExperimentOutput struct {
	ExperimentID string `json:"experiment_id,omitempty"`
	Result       string
}

func NewCreateNewExperimentUseCase(
	logger *slog.Logger,
	experimentRepository repository.ExperimentRepository) CreateNewExperimentUseCase {

	return CreateNewExperimentUseCase{logger: logger, experimentRepository: experimentRepository}
}

func (e CreateNewExperimentUseCase) Execute(
	cnei CreateNewExperimentInput) CreateNewExperimentOutput {

	_, exists, err := e.experimentRepository.GetByKey(cnei.Key)
	if exists {
		return CreateNewExperimentOutput{Result: "key already exists"}
	} else if err != nil {
		e.logger.Error("CreateNewExperimentUseCase#Execute",
			"error", err)
	}

	experiment := entity.Experiment{Key: cnei.Key, Name: cnei.Name}

	err = e.experimentRepository.Create(&experiment)
	if err != nil {
		e.logger.Error("CreateNewExperimentUseCase#Execute",
			"error", err)
	}

	return CreateNewExperimentOutput{ExperimentID: experiment.ID, Result: "Created successfully!"}
}
