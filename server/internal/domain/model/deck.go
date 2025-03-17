package model

import "github.com/google/uuid"

type DeckID string

func ParseDeckID(id string) (DeckID, error) {
	return DeckID(id), nil
}

func NewDeckID() DeckID {
	return DeckID(uuid.New().String())
}

func (id DeckID) String() string {
	return string(id)
}

type Deck struct {
	DeckID          DeckID         `dynamo:"DeckId,hash"`
	UserID          UserID         `dynamo:"UserId,index"`
	Name            string         `dynamo:"Name"`
	ThumbnailCardID MasterCardID   `dynamo:"ThumbnailCardId"`
	ThumbnailCard   *MasterCard    `dynamo:"-"`
	Color           MonsterType    `dynamo:"Color"`
	Energies        []MonsterType  `dynamo:"Energies"`
	MasterCardIDs   []MasterCardID `dynamo:"MasterCardIDs"`
	MasterCards     []*MasterCard  `dynamo:"-"`
}
