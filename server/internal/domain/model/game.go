package model

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

var ErrRoomNotFound = errors.New("room not found")
var ErrUserNotFound = errors.New("user not found")

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

type GamePhase string

const (
	GamePhaseReady        GamePhase = "ready"
	GamePhaseInitializing GamePhase = "initializing"
	GamePhaseBattle       GamePhase = "battle"
	GamePhaseEnd          GamePhase = "end"
)

type GameState struct {
	RoomID RoomID

	TurnPlayer    *Player
	NonTurnPlayer *Player

	Phase          GamePhase
	Turn           int
	LotteriedCoins []bool
}

func (m *GameState) IsMyTurn(userID UserID) bool {
	return m.TurnPlayer.UserID == userID
}

func (m *GameState) FindMeByUserID(userID UserID) (*Player, error) {
	if m.TurnPlayer.UserID == userID {
		return m.TurnPlayer, nil
	}

	if m.NonTurnPlayer.UserID == userID {
		return m.NonTurnPlayer, nil
	}

	return nil, ErrUserNotFound
}

func (m *GameState) FindEnemyByUserID(userID UserID) (*Player, error) {
	var enemy *Player
	if m.TurnPlayer.UserID == userID {
		enemy = m.NonTurnPlayer
	} else {
		enemy = m.TurnPlayer
	}

	if enemy == nil {
		return nil, xerrors.Errorf("enemy not found: %s", userID)
	}

	return enemy, nil
}

type Player struct {
	UserID          UserID
	BaseDeck        *Deck
	Deck            []*Card
	Hands           []*Card
	Fields          []*Card
	Trash           []*Card
	CurrentEnergies []MonsterType
	NextEnergy      MonsterType
	IsReady         bool

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
