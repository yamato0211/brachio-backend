package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type StoreUserInputPort interface {
	Execute(ctx context.Context, input *model.User) error
}

type StoreUserInteractor struct {
	UserRepository repository.UserRepository
}

func NewStoreUserUsecase(ur repository.UserRepository) StoreUserInputPort {
	return &StoreUserInteractor{
		UserRepository: ur,
	}
}

func (s *StoreUserInteractor) Execute(ctx context.Context, input *model.User) error {
	return s.UserRepository.Store(ctx, input)
}
