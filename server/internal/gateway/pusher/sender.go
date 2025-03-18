package websocket

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	wsmsg "github.com/yamato0211/brachio-backend/internal/handler/schema/websocket"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/websocket/event"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/websocket/payload"
	"github.com/yamato0211/brachio-backend/pkg/websocket"
	"google.golang.org/protobuf/proto"
)

type gameEventSender struct {
	pusher websocket.Pusher
}

func NewGameEventSender(pusher websocket.Pusher) service.GameEventSender {
	return &gameEventSender{
		pusher: pusher,
	}
}

func (g *gameEventSender) SendMatchingComplete(ctx context.Context, userID, oppoUserID model.UserID, roomID model.RoomID) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_MatchingCompleteEventToActor{
			MatchingCompleteEventToActor: &event.MatchingCompleteEventToActor{
				Payload: &payload.MatchingCompleteEventPayload{
					BattleId:   roomID.String(),
					OpponentId: oppoUserID.String(),
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, userID, b)
}

func (g *gameEventSender) SendDrawEffectEventToActor(ctx context.Context, actorID model.UserID, effects ...*messages.EffectWithSecret) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_DrawEffectEventToActor{
			DrawEffectEventToActor: &event.DrawEffectEventToActor{
				Payload: &payload.DrawEffectPayloadToActor{
					Effects: effects,
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, actorID, b)
}

func (g *gameEventSender) SendDrawEffectEventToRecipient(ctx context.Context, actorID model.UserID, effects ...*messages.Effect) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_DrawEffectEventToRecipient{
			DrawEffectEventToRecipient: &event.DrawEffectEventToRecipient{
				Payload: &payload.DrawEffectPayloadToRecipient{
					Effects: effects,
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, actorID, b)
}

func (g *gameEventSender) SendTurnStartEvent(ctx context.Context, userID model.UserID, turnUserID model.UserID) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_TurnStartEventToClients{
			TurnStartEventToClients: &event.TurnStartEventToClients{
				Payload: &payload.TurnStartPayload{
					UserId: turnUserID.String(),
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, userID, b)
}
