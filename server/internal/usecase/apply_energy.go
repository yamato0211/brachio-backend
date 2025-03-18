package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"golang.org/x/xerrors"
)

type ApplyEnergyInputPort interface {
	Execute(ctx context.Context, input *ApplyEnergyInput) error
}

type ApplyEnergyInput struct {
	RoomID  string
	UserID  string
	Postion int
}

type ApplyEnergyInteractor struct {
	GameStateRepository repository.GameStateRepository
}

func NewApplyEnergyUsecase(gameStateRepository repository.GameStateRepository) ApplyEnergyInputPort {
	return &ApplyEnergyInteractor{
		GameStateRepository: gameStateRepository,
	}
}

func (i *ApplyEnergyInteractor) Execute(ctx context.Context, input *ApplyEnergyInput) error {
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

		if state.TurnPlayer.UserID != userID {
			return xerrors.Errorf("not your turn")
		}

		if state.TurnPlayer.CurrentEnergy == nil {
			return xerrors.Errorf("energy is nil")
		}

		monster, err := state.TurnPlayer.GetMonsterByPosition(input.Postion)
		if err != nil {
			return err
		}

		monster.Energies = append(monster.Energies, *state.TurnPlayer.CurrentEnergy)

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	return nil
}
