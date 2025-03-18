package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
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
	GameEventSender     service.GameEventSender
}

func NewRetreatUsecase(
	gsr repository.GameStateRepository,
	ges service.GameEventSender,
) RetreatInputPort {
	return &RetreatInteractor{
		GameStateRepository: gsr,
		GameEventSender:     ges,
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

	var oppoID model.UserID

	var eventsForMe []*messages.EffectWithSecret
	var eventsForOppo []*messages.Effect
	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		if !state.IsMyTurn(userID) {
			return xerrors.Errorf("not your turn")
		}

		oppoID = state.NonTurnPlayer.UserID

		me, err := state.FindMeByUserID(userID)
		if err != nil {
			return err
		}

		if me.BenchMonsters[input.RetreatTo-1] == nil {
			return xerrors.Errorf("bench monster not found: %d", input.RetreatTo)
		}

		// Swap the battle monster and the bench monster
		me.BenchMonsters[input.RetreatTo-1], me.BattleMonster = me.BattleMonster, me.BenchMonsters[input.RetreatTo-1]

		eventsForMe = append(eventsForMe, i.makeEventForMe(input.RetreatTo))
		eventsForOppo = append(eventsForOppo, i.makeEventForOppo(input.RetreatTo))

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	if err := i.GameEventSender.SendDrawEffectEventToActor(ctx, userID, eventsForMe...); err != nil {
		return err
	}
	if err := i.GameEventSender.SendDrawEffectEventToRecipient(ctx, oppoID, eventsForOppo...); err != nil {
		return err
	}

	return nil
}

func (i *RetreatInteractor) makeEventForMe(position int) *messages.EffectWithSecret {
	return &messages.EffectWithSecret{
		Effect: &messages.EffectWithSecret_SwapBattleAndBench{
			SwapBattleAndBench: &messages.SwapBattleAndBenchEffect{
				Position: int32(position),
			},
		},
	}
}

func (i *RetreatInteractor) makeEventForOppo(position int) *messages.Effect {
	return &messages.Effect{
		Effect: &messages.Effect_SwapBattleAndBench{
			SwapBattleAndBench: &messages.SwapBattleAndBenchEffect{
				Position: int32(position),
			},
		},
	}
}
