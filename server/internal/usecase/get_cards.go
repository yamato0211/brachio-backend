package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMasterCardsInputPort interface {
	Execute(ctx context.Context) ([]*model.MasterCard, error)
}

type GetMasterCardsInteractor struct {
	MasterCardRepository repository.MasterCardRepository
}

func NewGetMasterCardsUsecase(mcr repository.MasterCardRepository) GetMasterCardsInputPort {
	return &GetMasterCardsInteractor{
		MasterCardRepository: mcr,
	}
}

func (g *GetMasterCardsInteractor) Execute(ctx context.Context) ([]*model.MasterCard, error) {
	return g.MasterCardRepository.FindAll(ctx)
}
