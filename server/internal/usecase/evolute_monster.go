package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
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
}

func NewEvoluteMonsterUsecase(
	gameStateRepository repository.GameStateRepository,
) EvoluteMonsterInputPort {
	return &EvoluteMonsterInteractor{
		GameStateRepository: gameStateRepository,
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

	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		if state.TurnPlayer.UserID != userID {
			return xerrors.Errorf("you are not turn player")
		}

		monster, err := state.TurnPlayer.GetMonsterByPosition(input.Position)
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

		return i.GameStateRepository.Store(ctx, state)
	})

	return err
}
