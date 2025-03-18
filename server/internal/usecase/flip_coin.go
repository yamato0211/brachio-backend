package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
)

type FlipCoinInputPort interface {
	Execute(ctx context.Context, input *FlipCoinInput) error
}

type FlipCoinInput struct {
	RoomID string
	UserID string
}

type FlipCoinInteractor struct {
	GameStateRepository repository.GameStateRepository
	GameEventSender     service.GameEventSender
}

func NewFlipCoinUsecase(
	gsr repository.GameStateRepository,
	ges service.GameEventSender,
) FlipCoinInputPort {
	return &FlipCoinInteractor{
		GameStateRepository: gsr,
		GameEventSender:     ges,
	}
}

func (i *FlipCoinInteractor) Execute(ctx context.Context, input *FlipCoinInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	state, err := i.GameStateRepository.Find(ctx, roomID)
	if err != nil {
		return err
	}

	if !state.IsMyTurn(userID) {
		return nil
	}

	event := i.makeEventForOppo(state.LotteriedCoins)
	if err := i.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, event); err != nil {
		return err
	}

	return nil
}

func (i *FlipCoinInteractor) makeEventForOppo(results []bool) *messages.Effect {
	return &messages.Effect{
		Effect: &messages.Effect_CoinToss{
			CoinToss: &messages.CoinTossEffect{
				Result: results,
			},
		},
	}
}
