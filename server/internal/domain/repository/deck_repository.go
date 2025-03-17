package repository

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type DeckRepository interface {
	Find(ctx context.Context, userID model.UserID, deckID model.DeckID) (*model.Deck, error)
	FindAll(ctx context.Context, userID model.UserID) ([]*model.Deck, error)
	Store(ctx context.Context, deck *model.Deck) error
}
