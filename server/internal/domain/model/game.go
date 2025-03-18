package model

import (
	"errors"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/xerrors"
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

func (m RoomID) String() string {
	return string(m)
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

func (m *Player) GetMonsterByPosition(position int) (*Monster, error) {
	if position == 0 {
		if m.BattleMonster == nil {
			return nil, xerrors.Errorf("position(%d) is invalid", position)
		}
		return m.BattleMonster, nil
	}

	if position < 0 || 3 < position {
		return nil, xerrors.Errorf("position(%d) is invalid", position)
	}

	monster := m.BenchMonsters[position-1]
	if monster == nil {
		return nil, xerrors.Errorf("position(%d) is invalid", position)
	}

	return monster, nil
}

func (m *Player) SetMonsterByPosition(position int, monster *Monster) error {
	if position == 0 {
		m.BattleMonster = monster
		return nil
	}

	if position < 0 || 3 < position {
		return xerrors.Errorf("position(%d) is invalid", position)
	}

	m.BenchMonsters[position-1] = monster
	return nil
}
