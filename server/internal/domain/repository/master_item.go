package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type MasterItemRepository interface {
	FindByMasterItemID(ctx context.Context, masterItemID model.MasterItemID) (*model.MasterItem, error)
	FindAll(ctx context.Context) ([]*model.MasterItem, error)
}
