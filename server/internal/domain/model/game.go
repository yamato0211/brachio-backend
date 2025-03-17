package model

import (
	"errors"

	"github.com/google/uuid"
)

var ErrRoomNotFound = errors.New("room not found")

type RoomID string

func NewRoomID() RoomID {
	id := uuid.New()
	return RoomID(id.String())
}

type GameState struct {
	RoomID RoomID

	Player1 *Player
	Player2 *Player

	Turn int
}

type Player struct {
	UserID   UserID
	BaseDeck *Deck
	Deck     []*Card
	Hands    []*Card
	Trash    []*Card
	Energies []Element

	Point         int
	BattleMonster *Monster
	BenchMonsters []*Monster
}
