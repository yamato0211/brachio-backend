package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetUserInputPort interface {
	Execute(ctx context.Context, userID string) (*model.User, error)
}

type GetUserInteractor struct {
	userRepository repository.UserRepository
}

func NewGetUserUsecase(ur repository.UserRepository) GetUserInputPort {
	return &GetUserInteractor{
		userRepository: ur,
	}
}

func (i *GetUserInteractor) Execute(ctx context.Context, userID string) (*model.User, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}

	user, err := i.userRepository.Find(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
