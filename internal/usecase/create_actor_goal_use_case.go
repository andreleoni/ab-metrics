package usecase

import (
	"ab-metrics/internal/domain/entity"
	"ab-metrics/internal/domain/repository"
	"log/slog"
)

type CreateActorGoalCheckUseCase struct {
	logger          *slog.Logger
	actorRepository repository.ActorRepository
	goalRepository  repository.GoalRepository
}

type CreateActorGoalCheckInput struct {
	GoalKey string
	ActorID string
}

type CreateActorGoalCheckOuput struct {
	Status string
}

func NewCreateActorGoalCheckUseCase(
	logger *slog.Logger,
	actorRepository repository.ActorRepository,
	goalRepository repository.GoalRepository) CreateActorGoalCheckUseCase {

	return CreateActorGoalCheckUseCase{logger: logger, actorRepository: actorRepository, goalRepository: goalRepository}
}

func (cagcuc CreateActorGoalCheckUseCase) Execute(
	cagcuci CreateActorGoalCheckInput) (CreateActorGoalCheckOuput, error) {

	_, err := cagcuc.goalRepository.Get(cagcuci.ActorID, cagcuci.GoalKey)
	if err != nil && err.Error() == "not found" {
		cagcuc.logger.Debug("CreateActorGoalCheckUseCase#Execute: not found error",
			"actor_id", cagcuci.ActorID,
			"goal_key", cagcuci.GoalKey)

		return CreateActorGoalCheckOuput{Status: "AlreadyAcomplished"}, nil
	} else if err != nil {
		cagcuc.logger.Error("CreateActorGoalCheckUseCase#Execute: error found on get actor goal",
			"actor_id", cagcuci.ActorID,
			"goal_key", cagcuci.GoalKey,
			"error", err.Error())
	}

	goal := entity.Goal{ActorID: cagcuci.ActorID, Key: cagcuci.GoalKey}

	err = cagcuc.goalRepository.Create(goal)
	if err != nil {
		cagcuc.logger.Error("error on creating goal",
			"goal", goal,
			"error", err.Error())
	}

	return CreateActorGoalCheckOuput{Status: "Accomplished"}, nil
}
