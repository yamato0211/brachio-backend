package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type DeleteMyDeckInputPort interface {
	Execute(ctx context.Context, deckID string) error
}

type DeleteMyDeckInteractor struct {
	deckRepository repository.DeckRepository
}

func NewDeleteMyDeckUsecase(dr repository.DeckRepository) DeleteMyDeckInputPort {
	return &DeleteMyDeckInteractor{
		deckRepository: dr,
	}
}

func (d *DeleteMyDeckInteractor) Execute(ctx context.Context, deckID string) error {
	did, err := model.ParseDeckID(deckID)
	if err != nil {
		return err
	}
	return d.deckRepository.Delete(ctx, did)
}
