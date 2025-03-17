package model

import "github.com/google/uuid"

type MasterCardID uuid.UUID

func NewMasterCardID() MasterCardID {
	return MasterCardID(uuid.New())
}

func ParseMasterCardID(s string) (MasterCardID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return MasterCardID{}, err
	}
	return MasterCardID(id), nil
}

func (id MasterCardID) String() string {
	return uuid.UUID(id).String()
}

type MasterCard struct {
	MasterCardID MasterCardID `dynamo:"MasterCardID,hash"`
	Name         string       `dynamo:"Name"`
	CardType     string       `dynamo:"CardType"` // モンスター,サポート　,グッズ
	Rarity       int          `dynamo:"Rarity"`   // レアリティ
	ImageURL     string       `dynamo:"ImageURL"`
	Expansion    string       `dynamo:"Expansion"` // 拡張パック名

	// Monster
	HP          int      `dynamo:"HP,omitempty"`
	Type        string   `dynamo:"Type,omitempty"`        // 属性 (e.g. fire)
	Weakness    string   `dynamo:"Weakness,omitempty"`    // 弱点 (e.g. water)
	Skills      []Skill  `dynamo:"Skills,omitempty"`      // ワザ
	Ability     *Ability `dynamo:"Ability,omitempty"`     // 特性
	RetreatCost int      `dynamo:"RetreatCost,omitempty"` // 逃げるコスト
	EvolvesFrom string   `dynamo:"EvolvesFrom,omitempty"` // 進化元
	EvelvesTo   string   `dynamo:"EvolvesTo,omitempty"`   // 進化先

	// Support & Goods
	Text string `dynamo:"Text,omitempty"` // 効果説明文
}

type Skill struct {
	Name         string   `dynamo:"Name"`
	Text         string   `dynamo:"Text,omitempty"` // 効果説明文 (e.g. コインを投げて表の場合20ダメージ追加)
	Damage       int      `dynamo:"Damage"`
	DamageOption string   `dynamo:"DamageOption,omitempty"` // ダメージオプション 20+の`+`の部分
	Cost         []string `dynamo:"Cost"`                   // 使用コスト (e.g. [fire, fire, water])
}

type Ability struct {
	Name string `dynamo:"Name"`
	Text string `dynamo:"Text"`
}
