package websocket

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/pkg/websocket"
)

type gameEventSender struct {
	pusher websocket.Pusher
}

func NewGameEventSender(pusher websocket.Pusher) service.GameEventSender {
	return &gameEventSender{
		pusher: pusher,
	}
}

// SendXXX implements service.GameEventSender.
func (g *gameEventSender) SendXXX(ctx context.Context, userID model.UserID, event any) error {
	panic("unimplemented")
}
