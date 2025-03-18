package model

type MasterCardID string

func NewMasterCardID(s string) MasterCardID {
	return MasterCardID(s)
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

type SubType string

const (
	MonsterSubTypeBasic  SubType = "basic"  // たね
	MonsterSubTypeStage1 SubType = "stage1" // 1進化
	MonsterSubTypeStage2 SubType = "stage2" // 2進化
)

type MasterCard struct {
	MasterCardID MasterCardID `dynamo:"MasterCardId,hash"`
	Name         string       `dynamo:"Name"`
	Description  string       `dynamo:"Description"`
	CardType     CardType     `dynamo:"CardType"` // モンスター,サポート　,グッズ
	Rarity       int          `dynamo:"Rarity"`   // レアリティ
	ImageURL     string       `dynamo:"ImageURL"`
	Expansion    string       `dynamo:"Expansion"` // 拡張パック名

	// Monster
	HP          int            `dynamo:"HP,omitempty"`
	SubType     SubType        `dynamo:"SubType,omitempty"`     // 進化段階
	Type        MonsterType    `dynamo:"Type,omitempty"`        // 属性 (e.g. fire)
	Weakness    MonsterType    `dynamo:"Weakness,omitempty"`    // 弱点 (e.g. water)
	Skills      []*Skill       `dynamo:"Skills,omitempty"`      // ワザ
	Ability     *Ability       `dynamo:"Ability,omitempty"`     // 特性
	RetreatCost int            `dynamo:"RetreatCost,omitempty"` // 逃げるコスト
	EvolvesFrom []MasterCardID `dynamo:"EvolvesFrom,omitempty"` // 進化元
	EvolvesTo   []MasterCardID `dynamo:"EvolvesTo,omitempty"`   // 進化先
	IsEx        bool           `dynamo:"IsEx,omitempty"`        // EXカード

	// Support & Goods
	Text string `dynamo:"Text,omitempty"` // 効果説明文
}

func (m *MasterCard) EvolvesFromSlice() []string {
	var evolvesFromStrings = make([]string, 0, len(m.EvolvesFrom))
	for _, e := range m.EvolvesFrom {
		evolvesFromStrings = append(evolvesFromStrings, string(e))
	}
	return evolvesFromStrings
}

func (m *MasterCard) EvelvesToSlice() []string {
	var evelvesToStrings = make([]string, 0, len(m.EvolvesTo))
	for _, e := range m.EvolvesTo {
		evelvesToStrings = append(evelvesToStrings, string(e))
	}
	return evelvesToStrings
}

type Skill struct {
	Name         string        `dynamo:"Name"`
	Text         string        `dynamo:"Text,omitempty"` // 効果説明文 (e.g. コインを投げて表の場合20ダメージ追加)
	Damage       int           `dynamo:"Damage"`
	DamageOption string        `dynamo:"DamageOption,omitempty"` // ダメージオプション 20+の`+`の部分
	Cost         []MonsterType `dynamo:"Cost"`                   // 使用コスト (e.g. [fire, fire, water])
}

func (s *Skill) CostSlice() []string {
	var costStrings = make([]string, 0, len(s.Cost))
	for _, c := range s.Cost {
		costStrings = append(costStrings, string(c))
	}
	return costStrings
}

type Ability struct {
	Name string `dynamo:"Name"`
	Text string `dynamo:"Text"`
}
