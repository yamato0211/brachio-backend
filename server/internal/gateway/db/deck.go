package db

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

const (
	deckTableName = "Decks"
	deckHashKey   = "DeckId"
	deckIndexKey  = "UserId"
)

type deckRepository struct {
	db dynamo.DB
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
	if err := d.db.Table(deckTableName).Get(deckIndexKey, userID).All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *deckRepository) Store(ctx context.Context, deck *model.Deck) error {
	return d.db.Table(masterCardsTableName).Put(deck).Run(ctx)
}

func NewDeckRepository(db dynamo.DB) repository.DeckRepository {
	return &deckRepository{db: db}
}
