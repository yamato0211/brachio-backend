package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type CreateMyDeckInputPort interface {
	Execute(ctx context.Context, userID string) (*model.Deck, error)
}

type CreateMyDeckInteractor struct {
	DeckRepository repository.DeckRepository
}

func (c *CreateMyDeckInteractor) Execute(ctx context.Context, userID string) (*model.Deck, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}
	deck := &model.Deck{
		DeckID: model.NewDeckID(),
		UserID: uid,
	}

	err = c.DeckRepository.Store(ctx, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func NewCreateMyDeckUsecase(dr repository.DeckRepository) CreateMyDeckInputPort {
	return &CreateMyDeckInteractor{
		DeckRepository: dr,
	}
}
