package model

type Monster struct {
	ID        string
	HP        int
	Energies  []MonsterType
	SpawnedAt int // 召喚されたターン
	BaseCard  *Card
	Knocked   bool // ノックアウトされているか

	RetreatCostAddition  int
	SkillCostAddition    int
	SkillDamageAddition  int
	SkillDamageReduction int
	AbilityUsed          bool
}

func (m *Monster) IsTypeEqual(t MonsterType) bool {
	return m.BaseCard.MasterCard.Type == t
}
