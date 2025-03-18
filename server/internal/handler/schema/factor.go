package schema

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

func MasterMonsterCardFromEntity(e *model.MasterCard) (*MasterMonsterCard, error) {
	if e == nil {
		return nil, fmt.Errorf("monster master card entity is nil")
	}
	return &MasterMonsterCard{
		Ability: &Ability{
			Name: e.Ability.Name,
			Text: e.Ability.Text,
		},
		CardType:     MasterCardType(e.CardType),
		Element:      Element(e.Type),
		EvolvesFrom:  lo.ToPtr(e.EvolvesFromSlice()),
		EvolvesTo:    lo.ToPtr(e.EvelvesToSlice()),
		Expansion:    lo.ToPtr(e.Expansion),
		Hp:           e.HP,
		ImageUrl:     e.ImageURL,
		IsEx:         &e.IsEx,
		MasterCardId: lo.ToPtr(e.MasterCardID.String()),
		Name:         e.Name,
		Rarity:       e.Rarity,
		RetreatCost:  &e.RetreatCost,
		Skills: lo.Map(e.Skills, func(skill model.Skill, _ int) Skill {
			return Skill{
				Cost: lo.Map(skill.Cost, func(cost model.MonsterType, _ int) Element {
					return Element(cost)
				}),
				Name:         skill.Name,
				Text:         skill.Text,
				Damage:       skill.Damage,
				DamageOption: lo.ToPtr(SkillDamageOption(skill.DamageOption)),
			}
		}),
		SubType:  lo.ToPtr(MasterMonsterCardSubType(e.SubType)),
		Weakness: Element(e.Weakness),
	}, nil
}

func MasterGoodsCardFromEntity(e *model.MasterCard) (*MasterGoodsCard, error) {
	if e == nil {
		return nil, fmt.Errorf("goods master card entity is nil")
	}

	return &MasterGoodsCard{
		CardType:     MasterCardType(e.CardType),
		Expansion:    lo.ToPtr(e.Expansion),
		ImageUrl:     e.ImageURL,
		MasterCardId: lo.ToPtr(e.MasterCardID.String()),
		Name:         e.Name,
		Rarity:       &e.Rarity,
		Text:         e.Text,
	}, nil
}

func MasterSupportCardFromEntity(e *model.MasterCard) (*MasterSupporterCard, error) {
	if e == nil {
		return nil, fmt.Errorf("support master card entity is nil")
	}

	return &MasterSupporterCard{
		CardType:     MasterCardType(e.CardType),
		Expansion:    lo.ToPtr(e.Expansion),
		ImageUrl:     e.ImageURL,
		MasterCardId: lo.ToPtr(e.MasterCardID.String()),
		Name:         e.Name,
		Rarity:       &e.Rarity,
		Text:         e.Text,
	}, nil
}

func FactoryCard(masterCard model.MasterCard) (*Card, error) {
	var card *Card
	switch masterCard.CardType {
	case model.CardTypeMonster:
		sc, err := MasterMonsterCardFromEntity(&masterCard)
		if err != nil {
			return nil, err
		}
		if err := card.MergeMasterMonsterCard(*sc); err != nil {
			return nil, err
		}
		return card, nil
	case model.CardTypeGoods:
		sc, err := MasterGoodsCardFromEntity(&masterCard)
		if err != nil {
			return nil, err
		}
		if err := card.MergeMasterGoodsCard(*sc); err != nil {
			return nil, err
		}
		return card, nil
	case model.CardTypeSupporter:
		sc, err := MasterSupportCardFromEntity(&masterCard)
		if err != nil {
			return nil, err
		}
		if err := card.MergeMasterSupporterCard(*sc); err != nil {
			return nil, err
		}
		return card, nil
	default:
		return nil, fmt.Errorf("unknown card type: %s", masterCard.CardType)
	}
}

func DeckWithIdFromEntity(e *model.Deck) (*DeckWithId, error) {
	if e == nil {
		return nil, fmt.Errorf("deck entity is nil")
	}

	sc, err := FactoryCard(*e.ThumbnailCard)
	if err != nil {
		return nil, err
	}

	myCards := make([]Card, 0, len(e.MasterCards))
	for _, mc := range e.MasterCards {
		fc, err := FactoryCard(*mc)
		if err != nil {
			return nil, err
		}
		myCards = append(myCards, *fc)
	}

	return &DeckWithId{
		Color: Element(e.Color),
		Energies: lo.Map(e.Energies, func(energy model.MonsterType, _ int) Element {
			return Element(energy)
		}),
		Cards:         myCards,
		Id:            lo.ToPtr(e.DeckID.String()),
		Name:          e.Name,
		ThumbnailCard: *sc,
	}, nil
}

func MasterCardWithFromEntity(e *model.MasterCard) (*MasterCard, error) {
	if e == nil {
		return nil, fmt.Errorf("master card entity is nil")
	}

	mc := &MasterCard{}
	switch e.CardType {
	case model.CardTypeMonster:
		sc, err := MasterMonsterCardFromEntity(e)
		if err != nil {
			return nil, err
		}
		if err := mc.MergeMasterMonsterCard(*sc); err != nil {
			return nil, err
		}
		return mc, nil
	case model.CardTypeGoods:
		sc, err := MasterGoodsCardFromEntity(e)
		if err != nil {
			return nil, err
		}
		if err := mc.MergeMasterGoodsCard(*sc); err != nil {
			return nil, err
		}
		return mc, nil
	case model.CardTypeSupporter:
		sc, err := MasterSupportCardFromEntity(e)
		if err != nil {
			return nil, err
		}
		if err := mc.MergeMasterSupporterCard(*sc); err != nil {
			return nil, err
		}
		return mc, nil
	default:
		return nil, fmt.Errorf("unknown card type: %s", e.CardType)
	}
}
