package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GiveUpInputPort interface {
	Execute(ctx context.Context, input *GiveUpInput) error
}

type GiveUpInput struct {
	RoomID string
	UserID string
}

type GiveUpInteractor struct {
	GameStateRepository repository.GameStateRepository
}

func NewGiveUpUsecase(gsr repository.GameStateRepository) GiveUpInputPort {
	return &GiveUpInteractor{
		GameStateRepository: gsr,
	}
}

func (i *GiveUpInteractor) Execute(ctx context.Context, input *GiveUpInput) error {
	// roomID, err := model.ParseRoomID(input.RoomID)
	// if err != nil {
	// 	return err
	// }

	// userID, err := model.ParseUserID(input.UserID)
	// if err != nil {
	// 	return err
	// }

	// err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
	// 	state, err := i.GameStateRepository.Find(ctx, roomID)
	// 	if err != nil {
	// 		return err
	// 	}

	// })
	return nil
}
