package usecase

import (
	"context"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	"golang.org/x/xerrors"
)

type AttackInputPort interface {
	Execute(ctx context.Context, input *AttackInput) error
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
	GameEventSender     service.GameEventSender
}

func NewAttackUsecase(
	gsr repository.GameStateRepository,
	sa service.SkillApprier,
	gm service.GameMasterService,
	ges service.GameEventSender,
) AttackInputPort {
	return &AttackInteractor{
		GameStateRepository: gsr,
		SkillApprier:        sa,
		GameMaster:          gm,
		GameEventSender:     ges,
	}
}

func (i *AttackInteractor) Execute(ctx context.Context, input *AttackInput) error {
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

		if !state.IsMyTurn(userID) {
			return xerrors.Errorf("not your turn")
		}

		me, err := state.FindMeByUserID(userID)
		if err != nil {
			return err
		}

		me.Effect, err = i.GameMaster.RunEffect(state, me.Effect, "before-attack", nil)
		if err != nil {
			return err
		}

		if _, err := i.SkillApprier.ApplySkill(ctx, state, input.SkillID); err != nil {
			return err
		}

		me.Effect, err = i.GameMaster.RunEffect(state, me.Effect, "after-attack", nil)
		if err != nil {
			return err
		}

		if err := i.GameEventSender.SendDrawEffectEventToActor(ctx, userID, i.makeEventForMe(0, state)); err != nil {
			return err
		}
		if err := i.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, i.makeEventForOppo(0, state)); err != nil {
			return err
		}

		isGameSet, winner := i.GameMaster.CheckWin(state)
		if isGameSet {
			winOrLose := &messages.DecideWinOrLoseEffect{UserId: winner.String()}
			if err := i.GameEventSender.SendDrawEffectEventToActor(ctx, userID, &messages.EffectWithSecret{Effect: &messages.EffectWithSecret_DecideWinOrLose{DecideWinOrLose: winOrLose}}); err != nil {
				return err
			}
			if err := i.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, &messages.Effect{Effect: &messages.Effect_DecideWinOrLose{DecideWinOrLose: winOrLose}}); err != nil {
				return err
			}
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

	if err := i.GameMaster.ChangeTurn(ctx, roomID); err != nil {
		return err
	}

	return nil
}

func (i *AttackInteractor) makeEventForMe(position int, state *model.GameState) *messages.EffectWithSecret {
	return &messages.EffectWithSecret{
		Effect: &messages.EffectWithSecret_Damage{
			Damage: &messages.DamageEffect{
				Position: int32(position),
				Amount:   lo.Map(state.Damages, func(d int, _ int) int32 { return int32(d) }),
			},
		},
	}
}

func (i *AttackInteractor) makeEventForOppo(position int, state *model.GameState) *messages.Effect {
	return &messages.Effect{
		Effect: &messages.Effect_Damage{
			Damage: &messages.DamageEffect{
				Position: int32(position),
				Amount:   lo.Map(state.Damages, func(d int, _ int) int32 { return int32(d) }),
			},
		},
	}
}
