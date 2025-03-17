package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type UserRepository interface {
	Find(ctx context.Context, userID model.UserID) (*model.User, error)
	Store(ctx context.Context, user *model.User) error
}
