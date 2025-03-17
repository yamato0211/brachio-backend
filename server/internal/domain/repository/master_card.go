package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type MasterCardRepository interface {
	FindByMasterCardID(ctx context.Context, masterCardID model.MasterCardID) (*model.MasterCard, error)
	FindAll(ctx context.Context) ([]*model.MasterCard, error)
}
