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

	myCards := make([]*model.MasterCard, 0, len(deck.MasterCardIDs))
	for _, cid := range deck.MasterCardIDs {
		mc, ok := lo.Find(masterCards, func(item *model.MasterCard) bool {
			return item.MasterCardID == cid
		})
		if !ok {
			return nil, errors.New("master card not found")
		}
		myCards = append(myCards, mc)
	}

	deck.MasterCards = myCards
	return deck, nil
}
