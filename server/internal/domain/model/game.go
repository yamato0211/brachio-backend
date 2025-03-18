package model

import (
	"errors"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

var ErrRoomNotFound = errors.New("room not found")

type RoomID string

func NewRoomID() RoomID {
	id := uuid.New()
	return RoomID(id.String())
}

func ParseRoomID(s string) (RoomID, error) {
	return RoomID(s), nil
}

type GameState struct {
	RoomID RoomID

	TurnPlayer    *Player
	NonTurnPlayer *Player

	Turn int
}

func (m *GameState) FindPlayerByUserID(userID UserID) *Player {
	return lo.Ternary(m.TurnPlayer.UserID == userID, m.TurnPlayer, m.NonTurnPlayer)
}

func (m *GameState) FindEnemyByUserID(userID UserID) *Player {
	return lo.Ternary(m.TurnPlayer.UserID == userID, m.NonTurnPlayer, m.TurnPlayer)
}

type Player struct {
	UserID        UserID
	BaseDeck      *Deck
	Deck          []*Card
	Hands         []*Card
	Fields        []*Card
	Trash         []*Card
	CurrentEnergy *MonsterType
	NextEnergy    *MonsterType

	Effect []*Effect

	Point         int
	BattleMonster *Monster
	BenchMonsters []*Monster
}

type Effect struct {
	ID      string
	Trigger string
	Fn      func(any) (bool, error)
	UserID  UserID
}
