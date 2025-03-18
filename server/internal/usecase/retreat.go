package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"golang.org/x/xerrors"
)

type RetreatInputPort interface {
	Execute(ctx context.Context, input *RetreatInput) error
}

type RetreatInput struct {
	RoomID    string
	UserID    string
	RetreatTo int
}

type RetreatInteractor struct {
	GameStateRepository repository.GameStateRepository
}

func NewRetreatUsecase(
	gsr repository.GameStateRepository,
) RetreatInputPort {
	return &RetreatInteractor{
		GameStateRepository: gsr,
	}
}

func (i *RetreatInteractor) Execute(ctx context.Context, input *RetreatInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		me := state.FindPlayerByUserID(userID)

		if me.BenchMonsters[input.RetreatTo-1] == nil {
			return xerrors.Errorf("bench monster not found: %d", input.RetreatTo)
		}

		// Swap the battle monster and the bench monster
		me.BenchMonsters[input.RetreatTo-1], me.BattleMonster = me.BattleMonster, me.BenchMonsters[input.RetreatTo-1]

		return i.GameStateRepository.Store(ctx, state)
	})

	return err
}
