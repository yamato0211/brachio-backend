package db

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

const (
	deckTableName   = "Decks"
	deckHashKey     = "DeckId"
	deckIndexKey    = "UserId"
	deckUserIdIndex = "UserIdIndex"
)

type deckRepository struct {
	db dynamo.DB
}

func (d *deckRepository) FindAllTempalte(ctx context.Context) ([]*model.Deck, error) {
	var data []*model.Deck
	if err := d.db.Table(deckTableName).Scan().Filter("begins_with(DeckId, ?)", "template-").All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *deckRepository) Find(ctx context.Context, deckID model.DeckID) (*model.Deck, error) {
	var data model.Deck
	if err := d.db.Table(deckTableName).Get(deckHashKey, deckID).One(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *deckRepository) FindAll(ctx context.Context, userID model.UserID) ([]*model.Deck, error) {
	var data []*model.Deck
	if err := d.db.Table(deckTableName).Get(deckIndexKey, userID).Index(deckUserIdIndex).All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *deckRepository) Store(ctx context.Context, deck *model.Deck) error {
	return d.db.Table(deckTableName).Put(deck).Run(ctx)
}

func (d *deckRepository) Delete(ctx context.Context, deckID model.DeckID) error {
	return d.db.Table(deckTableName).Delete(deckHashKey, deckID).Run(ctx)
}

func NewDeckRepository(db *dynamo.DB) repository.DeckRepository {
	return &deckRepository{db: *db}
}
