package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"golang.org/x/xerrors"
)

type ApplyAbilityInputPort interface {
	Execute(ctx context.Context, input *ApplyAbilityInput) error
}

type ApplyAbilityInput struct {
	RoomID   string
	UserID   string
	position int
}

type ApplyAbilityInteractor struct {
	GameStateRepository repository.GameStateRepository
	AbilityApplier      service.AbilityApplier
}

func NewApplyAbilityUsecase(
	gsr repository.GameStateRepository,
	aa service.AbilityApplier,
) ApplyAbilityInputPort {
	return &ApplyAbilityInteractor{
		GameStateRepository: gsr,
		AbilityApplier:      aa,
	}
}

func (i *ApplyAbilityInteractor) Execute(ctx context.Context, input *ApplyAbilityInput) error {
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

		var monster *model.Monster
		if input.position == 0 {
			monster = state.TurnPlayer.BattleMonster
		} else if 1 <= input.position && input.position <= 3 {
			monster = state.TurnPlayer.BenchMonsters[input.position-1]
		}

		if monster == nil {
			return xerrors.Errorf("monster not found")
		}

		if monster.BaseCard.MasterCard.Ability == nil {
			return xerrors.Errorf("monster has no ability")
		}

		return nil
	})

	return err
}
