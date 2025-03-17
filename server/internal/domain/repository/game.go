package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type GameStateRepository interface {
	Transaction(ctx context.Context, roomID model.RoomID, fn func(ctx context.Context) error) error
	Find(ctx context.Context, id model.RoomID) (*model.GameState, error)
	Save(ctx context.Context, state *model.GameState) error
	Delete(ctx context.Context, id model.RoomID) error
}
