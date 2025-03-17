package db

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

const (
	masterItemsTableName = "MasterItems"
	masterItemHashKey    = "MasterItemId"
)

type masterItemRepository struct {
	db *dynamo.DB
}

func (m *masterItemRepository) FindAll(ctx context.Context) ([]*model.MasterItem, error) {
	var data []*model.MasterItem
	if err := m.db.Table(masterItemsTableName).Scan().All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (m *masterItemRepository) FindByMasterItemID(ctx context.Context, masterItemID model.MasterItemID) (*model.MasterItem, error) {
	var data model.MasterItem
	if err := m.db.Table(masterItemsTableName).Get(masterItemHashKey, masterItemID).One(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func NewMasterItemRepository(db *dynamo.DB) repository.MasterItemRepository {
	return &masterItemRepository{db: db}
}
