package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type FindUserInputPort interface {
	Execute(ctx context.Context, userID string) (*model.User, error)
}

type FindUserInteractor struct {
	UserRepository repository.UserRepository
}

func NewFindUserUsecase(ur repository.UserRepository) FindUserInputPort {
	return &FindUserInteractor{
		UserRepository: ur,
	}
}

func (f *FindUserInteractor) Execute(ctx context.Context, userID string) (*model.User, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}
	return f.UserRepository.Find(ctx, uid)
}
