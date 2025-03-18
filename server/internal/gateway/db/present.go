package db

import (
	"context"
	"errors"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

var ErrPresentNotFound = errors.New("present not found")

const (
	presentTable    = "Presents"
	presentHashKey  = "PresentId"
	preesntRangeKey = "Time"
)

type presentRepository struct {
	db *dynamo.DB
}

func (p *presentRepository) Find(ctx context.Context, presentID model.PresentID) (*model.Present, error) {
	var present model.Present
	if err := p.db.Table(presentTable).Get(presentHashKey, presentID).One(ctx, &present); err != nil {
		return nil, err
	}
	return &present, nil
}

func (p *presentRepository) FindAll(ctx context.Context) ([]*model.Present, error) {
	var presents []*model.Present
	if err := p.db.Table(presentTable).Scan().All(ctx, &presents); err != nil {
		return nil, err
	}
	return presents, nil
}

func (p *presentRepository) Store(ctx context.Context, present *model.Present) error {
	return p.db.Table(presentTable).Put(present).Run(ctx)
}

func NewPresentRepository(db *dynamo.DB) repository.PresentRepository {
	return &presentRepository{db: db}
}
