package usecase

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/internal/domain/repository"
	"log/slog"
)

type CreateActorGoalCheckUseCase struct {
	logger         *slog.Logger
	goalRepository repository.GoalRepository
}

type CreateActorGoalCheckInput struct {
	GoalKey string
	ActorID string
}

type CreateActorGoalCheckOuput struct {
	Status string
}

func NewCreateActorGoalCheckUseCase(logger *slog.Logger) CreateActorGoalCheckUseCase {
	return CreateActorGoalCheckUseCase{logger: logger}
}

func (cagcuc CreateActorGoalCheckUseCase) Execute(
	cagcuci CreateActorGoalCheckInput) (CreateActorGoalCheckOuput, error) {

	_, goalExists := cagcuc.goalRepository.Get(cagcuci.ActorID, cagcuci.GoalKey)
	if goalExists {
		cagcuc.logger.Debug("CreateActorGoalCheckUseCase#Execute",
			"actor_id", cagcuci.ActorID,
			"goal_key", cagcuci.GoalKey)

		return CreateActorGoalCheckOuput{Status: "AlreadyAcomplished"}, nil
	}

	goal := entity.Goal{ActorID: cagcuci.ActorID, Key: cagcuci.GoalKey}

	goal, err := cagcuc.goalRepository.Create(goal)
	if err != nil {
		cagcuc.logger.Error("error on creating goal",
			"goal", goal,
			"error", err.Error())
	}

	return CreateActorGoalCheckOuput{Status: "Accomplished"}, nil
}
