package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type PresentRepository interface {
	Find(ctx context.Context, presentID model.PresentID) (*model.Present, error)
	FindAll(ctx context.Context) ([]*model.Present, error)
	Store(ctx context.Context, present *model.Present) error
}
