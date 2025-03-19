package websocket

import (
	"context"

	"github.com/samber/lo"
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
	pusher *websocket.Pusher
}

func NewGameEventSender(pusher *websocket.Pusher) service.GameEventSender {
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

func (g *gameEventSender) SendDrawCardsEventToActor(ctx context.Context, actorID model.UserID, deckCount int, cards ...*model.Card) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_DrawEventToActor{
			DrawEventToActor: &event.DrawEventToActor{
				Payload: &payload.DrawCardIndividualPayload{
					Count:  int32(len(cards)),
					Cards:  lo.Map(cards, func(card *model.Card, _ int) *messages.Card { return messages.NewCard(card) }),
					Remain: int32(deckCount),
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

func (g *gameEventSender) SendDrawCardsEventToRecipient(ctx context.Context, actorID model.UserID, deckCount int, cards ...*model.Card) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_DrawEventToRecipient{
			DrawEventToRecipient: &event.DrawEventToRecipient{
				Payload: &payload.DrawCardPayload{
					Count:  int32(len(cards)),
					Remain: int32(deckCount),
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

func (g *gameEventSender) SendNextEnergyEventToActor(ctx context.Context, actorID model.UserID, energy model.MonsterType) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_NextEnergyEventToActor{
			NextEnergyEventToActor: &event.NextEnergyEventToActor{
				Payload: &payload.NextEnergyPayload{
					Energy: messages.NewElement(energy),
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

func (g *gameEventSender) SendNextEnergyEventToRecipient(ctx context.Context, opponentID model.UserID, energy model.MonsterType) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_NextEnergyEventToActor{
			NextEnergyEventToActor: &event.NextEnergyEventToActor{
				Payload: &payload.NextEnergyPayload{
					Energy: messages.NewElement(energy),
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, opponentID, b)
}

func (g *gameEventSender) SendDecideOrderEvent(ctx context.Context, userID, firstUserID, secondUserID model.UserID) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_DecideOrderEventToActor{
			DecideOrderEventToActor: &event.DecideOrderEventToActor{
				Payload: &payload.DecideOrderEventPayload{
					FirstUserId:  firstUserID.String(),
					SecondUserId: secondUserID.String(),
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

func (g *gameEventSender) SendGameStartEvent(ctx context.Context, userID model.UserID) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_StartGameEventToClients{
			StartGameEventToClients: &event.StartGameEventToClients{
				Payload: &payload.StartGamePayload{},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, userID, b)
}
