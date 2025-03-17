package model

type DeckID string

func ParseDeckID(id string) (DeckID, error) {
	return DeckID(id), nil
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
