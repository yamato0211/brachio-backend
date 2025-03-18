package db

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

const (
	masterCardsTableName = "MasterCards"
	masterCardsHashKey   = "MasterCardId"
)

type masterCardRepository struct {
	db dynamo.DB
}

func NewMasterCardRepository(db *dynamo.DB) repository.MasterCardRepository {
	return &masterCardRepository{db: *db}
}

func (r *masterCardRepository) FindAll(ctx context.Context) ([]*model.MasterCard, error) {
	var data []*model.MasterCard
	if err := r.db.Table(masterCardsTableName).Scan().All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *masterCardRepository) FindByMasterCardID(ctx context.Context, id model.MasterCardID) (*model.MasterCard, error) {
	var data model.MasterCard
	if err := r.db.Table(masterCardsTableName).Get(masterCardsHashKey, id).One(ctx, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *masterCardRepository) Store(ctx context.Context, masterCard *model.MasterCard) error {
	return r.db.Table(masterCardsTableName).Put(masterCard).Run(ctx)
}

func (r *masterCardRepository) StoreAll(ctx context.Context, masterCards []*model.MasterCard) error {
	_, err := r.db.Table(masterCardsTableName).Batch().Write().Put(masterCards).Run(ctx)
	return err
}
