package service

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
)

type GameEventSender interface {
	SendMatchingComplete(ctx context.Context, userID, oppoUserID model.UserID, roomID model.RoomID) error
	SendGameStartEvent(ctx context.Context, userID model.UserID) error
	SendTurnStartEvent(ctx context.Context, userID, turnUserID model.UserID) error
	SendDrawCardsEventToActor(ctx context.Context, actorID model.UserID, deckCount int, cards ...*model.Card) error
	SendDrawCardsEventToRecipient(ctx context.Context, opponentID model.UserID, deckCount int, cards ...*model.Card) error
	SendNextEnergyEventToActor(ctx context.Context, actorID model.UserID, energy model.MonsterType) error
	SendNextEnergyEventToRecipient(ctx context.Context, opponentID model.UserID, energy model.MonsterType) error
	SendDecideOrderEvent(ctx context.Context, userID, firstUserID, secondUserID model.UserID) error

	SendDrawEffectEventToActor(ctx context.Context, actorID model.UserID, effects ...*messages.EffectWithSecret) error
	SendDrawEffectEventToRecipient(ctx context.Context, actorID model.UserID, effects ...*messages.Effect) error
}
