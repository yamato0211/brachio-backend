package model

import (
	"errors"

	"github.com/google/uuid"
)

var ErrItemNotFound = errors.New("item not found")

type MasterItemID string

func NewMasterItemID() MasterItemID {
	return MasterItemID(uuid.New().String())
}

func ParseMasterItemID(s string) (MasterItemID, error) {
	return MasterItemID(s), nil
}

func (id MasterItemID) String() string {
	return string(id)
}

type MasterItem struct {
	ItemID      MasterItemID `dynamo:"MasterItemId,hash"`
	Name        string       `dynamo:"Name"`
	Description string       `dynamo:"Description"`
	ImageURL    string       `dynamo:"ImageUrl"`
}
