package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type UpdateMyDeckInputPort interface {
	Execute(ctx context.Context, input *model.Deck) error
}

type UpdateMyDeckInteractor struct {
	DeckRepository repository.DeckRepository
}

func (u *UpdateMyDeckInteractor) Execute(ctx context.Context, input *model.Deck) error {
	return u.DeckRepository.Store(ctx, input)
}

func NewUpdateMyDeckUsecase(dr repository.DeckRepository) UpdateMyDeckInputPort {
	return &UpdateMyDeckInteractor{
		DeckRepository: dr,
	}
}
