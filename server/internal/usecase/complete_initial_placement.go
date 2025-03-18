package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
)

type CompleteInitialPlacementInputPort interface {
	Execute(ctx context.Context, input *CompleteInitialPlacementInput) error
}

type CompleteInitialPlacementInput struct {
	RoomID string
	UserID string
}

type CompleteInitialPlacementInteractor struct {
	GameMaster service.GameMasterService
}

func NewCompleteInitialPlacementUsecase(
	gm service.GameMasterService,
) CompleteInitialPlacementInputPort {
	return &CompleteInitialPlacementInteractor{
		GameMaster: gm,
	}
}

func (i *CompleteInitialPlacementInteractor) Execute(ctx context.Context, input *CompleteInitialPlacementInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	return i.GameMaster.ReadyForStart(ctx, roomID, userID)
}
