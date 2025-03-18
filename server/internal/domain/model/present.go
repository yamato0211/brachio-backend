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
	PresentID PresentID `dynamo:"PresentId,hash"`
}
