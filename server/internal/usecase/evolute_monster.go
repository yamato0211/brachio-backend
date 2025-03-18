package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	"golang.org/x/xerrors"
)

type EvoluteMonsterInputPort interface {
	Execute(ctx context.Context, input *EvoluteMonsterInput) error
}

type EvoluteMonsterInput struct {
	RoomID   string
	UserID   string
	CardID   string
	Position int
}

type EvoluteMonsterInteractor struct {
	GameStateRepository repository.GameStateRepository
	GameEventSender     service.GameEventSender
}

func NewEvoluteMonsterUsecase(
	gsr repository.GameStateRepository,
	ges service.GameEventSender,
) EvoluteMonsterInputPort {
	return &EvoluteMonsterInteractor{
		GameStateRepository: gsr,
		GameEventSender:     ges,
	}
}

// Execute implements EvoluteMonsterInputPort.
func (i *EvoluteMonsterInteractor) Execute(ctx context.Context, input *EvoluteMonsterInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	cardID, err := model.ParseCardID(input.CardID)
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

		monster, err := me.GetMonsterByPosition(input.Position)
		if err != nil {
			return err
		}

		card, idx, isFound := lo.FindIndexOf(state.TurnPlayer.Hands, func(c *model.Card) bool { return c.CardID == cardID })
		if !isFound {
			return xerrors.Errorf("card not found: %d", cardID)
		}

		if card.MasterCard.CardType != model.CardTypeMonster {
			return xerrors.Errorf("card is not monster: %d", cardID)
		}

		monster, err = card.Evolute(state.Turn, monster)
		if err != nil {
			return err
		}

		if err := state.TurnPlayer.SetMonsterByPosition(input.Position, monster); err != nil {
			return err
		}

		state.TurnPlayer.Hands = slices.Delete(state.TurnPlayer.Hands, idx, idx+1)

		eventsForMe = append(eventsForMe, i.makeEventForMe(input.Position, card))
		eventsForOppo = append(eventsForOppo, i.makeEventForOppo(input.Position, card))

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

	return err
}

func (i *EvoluteMonsterInteractor) makeEventForMe(position int, card *model.Card) *messages.EffectWithSecret {
	return &messages.EffectWithSecret{
		Effect: &messages.EffectWithSecret_Evolution{
			Evolution: &messages.EvolutionEffect{
				Card:     messages.NewCard(card),
				Position: int32(position),
			},
		},
	}
}

func (i *EvoluteMonsterInteractor) makeEventForOppo(position int, card *model.Card) *messages.Effect {
	return &messages.Effect{
		Effect: &messages.Effect_Evolution{
			Evolution: &messages.EvolutionEffect{
				Card:     messages.NewCard(card),
				Position: int32(position),
			},
		},
	}
}
