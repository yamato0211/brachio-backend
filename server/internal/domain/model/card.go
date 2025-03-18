package model

import (
	"slices"
	"strconv"

	"golang.org/x/xerrors"
)

type CardID int

func NewCardID(x int) CardID {
	return CardID(x)
}

func ParseCardID(s string) (CardID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return NewCardID(id), nil
}

func (i CardID) String() string {
	return strconv.Itoa(int(i))
}

type Card struct {
	CardID            CardID
	MasterCard        *MasterCard
	ReservedMonsterID string
}

func (m *Card) Summon(turn int) (*Monster, error) {
	if m.MasterCard.CardType != CardTypeMonster {
		return nil, xerrors.Errorf("card is not monster: %d", m.CardID)
	}

	return &Monster{
		ID:        m.ReservedMonsterID,
		HP:        m.MasterCard.HP,
		Energies:  make([]MonsterType, 0),
		BaseCard:  m,
		SpawnedAt: turn,
	}, nil
}

func (m *Card) Evolute(turn int, monster *Monster) (*Monster, error) {
	if monster.SpawnedAt == turn {
		return nil, xerrors.Errorf("monster is already summoned: %s", monster.ID)
	}

	if slices.Contains(monster.BaseCard.MasterCard.EvolvesTo, m.MasterCard.MasterCardID) {
		return nil, xerrors.Errorf("card is not evoluted card: %d", m.CardID)
	}

	evoluteFrom := monster.BaseCard.MasterCard
	monster.HP += m.MasterCard.HP - evoluteFrom.HP
	monster.SpawnedAt = turn
	monster.BaseCard = m

	return monster, nil
}
