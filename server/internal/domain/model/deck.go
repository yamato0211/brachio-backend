package model

type DeckID string

func ParseDeckID(id string) (DeckID, error) {
	return DeckID(id), nil
}

type Deck struct {
	ID DeckID
}
