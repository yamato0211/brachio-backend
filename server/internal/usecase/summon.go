package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type SummonInputPort interface {
	Execute(ctx context.Context, input *SummonInput) error
}

type SummonInput struct {
	RoomID   string
	UserID   string
	CardID   string
	Position int
}

type SummonInteractor struct {
	GameStateRepository repository.GameStateRepository
}

func NewSummonUsecase(
	gsr repository.GameStateRepository,
) SummonInputPort {
	return &SummonInteractor{
		GameStateRepository: gsr,
	}
}

func (i *SummonInteractor) Execute(ctx context.Context, input *SummonInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}
	_ = userID

	i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		return nil
	})
}
