package usecase

import (
	"context"
	"errors"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyDeckInputPort interface {
	Execute(ctx context.Context, deckID model.DeckID) (*model.Deck, error)
}

type GetMyDeckInteractor struct {
	DeckRepository       repository.DeckRepository
	MasterCardRepository repository.MasterCardRepository
}

func NewGetMyDeckUsecase(dr repository.DeckRepository, mr repository.MasterCardRepository) GetMyDeckInputPort {
	return &GetMyDeckInteractor{
		DeckRepository:       dr,
		MasterCardRepository: mr,
	}
}

func (g *GetMyDeckInteractor) Execute(ctx context.Context, deckID model.DeckID) (*model.Deck, error) {
	deck, err := g.DeckRepository.Find(ctx, deckID)
	if err != nil {
		return nil, err
	}

	masterCards, err := g.MasterCardRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	tc, ok := lo.Find(masterCards, func(item *model.MasterCard) bool {
		return item.MasterCardID == deck.ThumbnailCardID
	})
	if !ok {
		return nil, errors.New("thumbnail card not found")
	}
	deck.ThumbnailCard = tc

	myMasterCards := lo.Filter(masterCards, func(card *model.MasterCard, _ int) bool {
		for _, id := range deck.MasterCardIDs {
			if card.MasterCardID == id {
				return true
			}
		}
		return false
	})

	deck.MasterCards = myMasterCards
	return deck, nil
}
