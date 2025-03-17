package model

import "github.com/google/uuid"

type MasterCardID string

func NewMasterCardID() MasterCardID {
	return MasterCardID(uuid.New().String())
}

func ParseMasterCardID(s string) (MasterCardID, error) {
	return MasterCardID(s), nil
}

func (id MasterCardID) String() string {
	return string(id)
}

type CardType string

const (
	CardTypeMonster   CardType = "monster"
	CardTypeSupporter CardType = "supporter"
	CardTypeGoods     CardType = "goods"
)

type MonsterType string

const (
	MonsterTypeNull       MonsterType = "null"
	MonsterTypeMoney      MonsterType = "money"
	MonsterTypeKnowledge  MonsterType = "knowledge"
	MonsterTypeMuscle     MonsterType = "muscle"
	MonsterTypeAlchohol   MonsterType = "alchohol"
	MonsterTypePopularity MonsterType = "popularity"
)

type MasterCard struct {
	MasterCardID MasterCardID `dynamo:"CardId,hash"`
	Name         string       `dynamo:"Name"`
	Description  string       `dynamo:"Description"`
	CardType     CardType     `dynamo:"CardType"` // モンスター,サポート　,グッズ
	Rarity       int          `dynamo:"Rarity"`   // レアリティ
	ImageURL     string       `dynamo:"ImageURL"`
	Expansion    string       `dynamo:"Expansion"` // 拡張パック名

	// Monster
	HP          int            `dynamo:"HP,omitempty"`
	Type        MonsterType    `dynamo:"Type,omitempty"`        // 属性 (e.g. fire)
	Weakness    MonsterType    `dynamo:"Weakness,omitempty"`    // 弱点 (e.g. water)
	Skills      []Skill        `dynamo:"Skills,omitempty"`      // ワザ
	Ability     *Ability       `dynamo:"Ability,omitempty"`     // 特性
	RetreatCost int            `dynamo:"RetreatCost,omitempty"` // 逃げるコスト
	EvolvesFrom []MasterCardID `dynamo:"EvolvesFrom,omitempty"` // 進化元
	EvelvesTo   []MasterCardID `dynamo:"EvolvesTo,omitempty"`   // 進化先
	IsEx        bool           `dynamo:"IsEx,omitempty"`        // EXカード

	// Support & Goods
	Text string `dynamo:"Text,omitempty"` // 効果説明文
}

type Skill struct {
	Name         string        `dynamo:"Name"`
	Text         string        `dynamo:"Text,omitempty"` // 効果説明文 (e.g. コインを投げて表の場合20ダメージ追加)
	Damage       int           `dynamo:"Damage"`
	DamageOption string        `dynamo:"DamageOption,omitempty"` // ダメージオプション 20+の`+`の部分
	Cost         []MonsterType `dynamo:"Cost"`                   // 使用コスト (e.g. [fire, fire, water])
}

type Ability struct {
	Name string `dynamo:"Name"`
	Text string `dynamo:"Text"`
}
