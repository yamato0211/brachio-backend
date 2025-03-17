package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type MatchingRepository interface {
	Find(ctx context.Context, password string) (model.UserID, error)
	Store(ctx context.Context, userID model.UserID) error
	Delete(ctx context.Context, userID model.UserID) error
}
