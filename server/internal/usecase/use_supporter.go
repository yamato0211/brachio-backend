package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"golang.org/x/xerrors"
)

type UseSupporterInputPort interface {
	Execute(ctx context.Context, input *UseSupporterInput) error
}

type UseSupporterInput struct {
	RoomID  string
	UserID  string
	CardID  string
	Targets []int
}

type UseSupporterInteractor struct {
	GameStateRepository repository.GameStateRepository
	SupporterApplier    service.SupporterApplier
}

func NewUseSupporterUsecase(
	gsr repository.GameStateRepository,
	sa service.SupporterApplier,
) UseSupporterInputPort {
	return &UseSupporterInteractor{
		GameStateRepository: gsr,
		SupporterApplier:    sa,
	}
}

func (i *UseSupporterInteractor) Execute(ctx context.Context, input *UseSupporterInput) error {
	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	cardID, err := model.ParseCardID(input.CardID)
	if err != nil {
		return nil
	}

	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		if state.TurnPlayer.UserID != userID {
			return model.ErrRoomNotFound
		}

		card, idx, isFound := lo.FindIndexOf(state.TurnPlayer.Hands, func(card *model.Card) bool {
			return card.CardID == cardID
		})
		if !isFound {
			return xerrors.New("card not found")
		}
		state.TurnPlayer.Hands = slices.Delete(state.TurnPlayer.Hands, idx, idx+1)

		if err := i.SupporterApplier.ApplySupporter(state, card.MasterCard.MasterCardID, input.Targets); err != nil {
			return err
		}

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	return nil
}
