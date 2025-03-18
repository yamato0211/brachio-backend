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

type UseGoodsInputPort interface {
	Execute(ctx context.Context, input *UseGoodsInput) error
}

type UseGoodsInput struct {
	RoomID  string
	UserID  string
	CardID  string
	Targets []int
}

type UseGoodsInteractor struct {
	GameStateRepository repository.GameStateRepository
	GoodsApplier        service.GoodsApplier
}

func NewUserGoodsUsecase(
	gsr repository.GameStateRepository,
	ga service.GoodsApplier,
) UseGoodsInputPort {
	return &UseGoodsInteractor{
		GameStateRepository: gsr,
		GoodsApplier:        ga,
	}
}

func (i *UseGoodsInteractor) Execute(ctx context.Context, input *UseGoodsInput) error {
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

		if err := i.GoodsApplier.ApplyGoods(state, card.MasterCard.MasterCardID, input.Targets); err != nil {
			return err
		}

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	return nil
}
