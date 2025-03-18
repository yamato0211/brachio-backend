package model

import "github.com/google/uuid"

type PresentID string

func ParsePresentID(id string) (PresentID, error) {
	return PresentID(id), nil
}

func NewPresentID() PresentID {
	return PresentID(uuid.New().String())
}

func (id PresentID) String() string {
	return string(id)
}

type Present struct {
	PresentID       PresentID    `dynamo:"PresentId,hash"`
	Time            int          `dynamo:"Time"`
	Message         string       `dynamo:"Message"`
	ReceivedUserIDs []UserID     `dynamo:"ReceivedUserIds"`
	ItemID          MasterItemID `dynamo:"ItemId"`
	Item            MasterItem   `dynamo:"-"`
	ItemCount       int          `dynamo:"ItemCount"`
}
