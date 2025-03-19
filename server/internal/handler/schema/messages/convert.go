package messages

import (
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

func NewElement(m model.MonsterType) Element {
	switch m {
	case model.MonsterTypeAlchohol:
		return Element_ALCHOHOL
	case model.MonsterTypeKnowledge:
		return Element_KNOWLEDGE
	case model.MonsterTypePopularity:
		return Element_POPULARITY
	case model.MonsterTypeMoney:
		return Element_MONEY
	case model.MonsterTypeMuscle:
		return Element_MUSCLE
	case model.MonsterTypeNull:
		return Element_NULL
	default:
		return Element_ELEMENT_UNKNOWN
	}
}

func NewSubType(m model.SubType) SubType {
	switch m {
	case model.MonsterSubTypeBasic:
		return SubType_BASIC
	case model.MonsterSubTypeStage1:
		return SubType_STAGE1
	case model.MonsterSubTypeStage2:
		return SubType_STAGE2
	default:
		return SubType_SUB_TYPE_UNSPECIFIED
	}
}

func NewDamageOption(opt string) *DamageOption {
	switch opt {
	case "+":
		return lo.ToPtr(DamageOption_PLUS)
	case "x":
		return lo.ToPtr(DamageOption_X)
	default:
		return nil
	}
}

func NewSkill(m *model.Skill) *Skill {
	return &Skill{
		Name:         m.Name,
		Text:         m.Text,
		Damage:       int32(m.Damage),
		DamageOption: NewDamageOption(m.DamageOption),
		Cost:         lo.Map(m.Cost, func(e model.MonsterType, _ int) Element { return NewElement(e) }),
	}
}

func NewAbility(m *model.Ability) *Ability {
	if m == nil {
		return nil
	}

	return &Ability{
		Name: m.Name,
		Text: m.Text,
	}
}

func NewCardType(m model.CardType) MasterCardType {
	switch m {
	case model.CardTypeMonster:
		return MasterCardType_MONSTER
	case model.CardTypeSupporter:
		return MasterCardType_SUPPORTER
	case model.CardTypeGoods:
		return MasterCardType_GOODS
	default:
		return MasterCardType_MASTER_CARD_TYPE_UNSPECIFIED
	}
}

func NewCard(m *model.Card) *Card {
	return &Card{
		Id:         m.CardID.String(),
		MasterCard: NewMasterCard(m),
	}
}

func NewMasterCard(m *model.Card) *MasterCard {
	base := &MasterCardBase{
		MasterCardId: m.MasterCard.MasterCardID.String(),
		Name:         m.MasterCard.Name,
		CardType:     NewCardType(m.MasterCard.CardType),
		Rarity:       int32(m.MasterCard.Rarity),
		ImageUrl:     m.MasterCard.ImageURL,
		Expansion:    m.MasterCard.Expansion,
	}

	switch m.MasterCard.CardType {
	case model.CardTypeMonster:
		return &MasterCard{

			CardVariant: &MasterCard_MasterMonsterCard{
				MasterMonsterCard: &MasterMonsterCard{
					Base:        base,
					Hp:          int32(m.MasterCard.HP),
					SubType:     NewSubType(m.MasterCard.SubType),
					Element:     NewElement(m.MasterCard.Type),
					Weakness:    NewElement(m.MasterCard.Weakness),
					Skills:      lo.Map(m.MasterCard.Skills, func(s *model.Skill, _ int) *Skill { return NewSkill(s) }),
					Ability:     NewAbility(m.MasterCard.Ability),
					RetreatCost: int32(m.MasterCard.RetreatCost),
					EvolvesFrom: lo.Map(m.MasterCard.EvolvesFrom, func(id model.MasterCardID, _ int) string { return id.String() }),
					EvolvesTo:   lo.Map(m.MasterCard.EvolvesTo, func(id model.MasterCardID, _ int) string { return id.String() }),
				},
			},
		}
	case model.CardTypeSupporter:
		return &MasterCard{
			CardVariant: &MasterCard_MasterSupporterCard{
				MasterSupporterCard: &MasterSupporterCard{
					Base: base,
					Text: m.MasterCard.Text,
				},
			},
		}
	case model.CardTypeGoods:
		return &MasterCard{
			CardVariant: &MasterCard_MasterGoodsCard{
				MasterGoodsCard: &MasterGoodsCard{
					Base: base,
					Text: m.MasterCard.Text,
				},
			},
		}
	default:
		return nil
	}
}
