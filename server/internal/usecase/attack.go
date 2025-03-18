package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"golang.org/x/xerrors"
)

type AttackInputPort interface {
	Execute(ctx context.Context, input AttackInput) error
}

type AttackInput struct {
	RoomID  string
	UserID  string
	SkillID int
}

type AttackInteractor struct {
	GameStateRepository repository.GameStateRepository
	SkillApprier        service.SkillApprier
	GameMaster          service.GameMasterService
}

func NewAttackUsecase(
	gsr repository.GameStateRepository,
	sa service.SkillApprier,
	gm service.GameMasterService,
) AttackInputPort {
	return &AttackInteractor{
		GameStateRepository: gsr,
		SkillApprier:        sa,
		GameMaster:          gm,
	}
}

func (i *AttackInteractor) Execute(ctx context.Context, input AttackInput) error {
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
		if me == nil {
			return xerrors.Errorf("player not found: %s", userID)
		}

		enemy := state.FindEnemyByUserID(userID)
		if enemy == nil {
			return xerrors.Errorf("enemy not found: %s", userID)
		}

		me.Effect, err = i.GameMaster.RunEffect(me.Effect, "before-attack", nil)
		if err != nil {
			return err
		}

		damage, err := i.SkillApprier.ApplySkill(me, enemy, input.SkillID)
		if err != nil {
			return err
		}
		_ = damage

		me.Effect, err = i.GameMaster.RunEffect(me.Effect, "after-attack", nil)
		if err != nil {
			return err
		}

		// ゲームの状態を保存する
		err = i.GameStateRepository.Store(ctx, state)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
