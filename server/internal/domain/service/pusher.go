package service

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
)

type GameEventSender interface {
	SendMatchingComplete(ctx context.Context, userID, oppoUserID model.UserID, roomID model.RoomID) error
	SendDrawEffectEventToActor(ctx context.Context, actorID model.UserID, effects ...*messages.EffectWithSecret) error
	SendDrawEffectEventToRecipient(ctx context.Context, actorID model.UserID, effects ...*messages.Effect) error
	SendTurnStartEvent(ctx context.Context, userID, turnUserID model.UserID) error
	// Send
}
