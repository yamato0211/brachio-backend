package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyDecksInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.Deck, error)
}

type GetMyDecksInteractor struct {
	DeckRepository repository.DeckRepository
}

func NewGetMyDecksUsecase(dr repository.DeckRepository) GetMyDecksInputPort {
	return &GetMyDecksInteractor{
		DeckRepository: dr,
	}
}

func (g *GetMyDecksInteractor) Execute(ctx context.Context, userID string) ([]*model.Deck, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}
	return g.DeckRepository.FindAll(ctx, uid)
}
