package service

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"golang.org/x/xerrors"
)

type GoodsApplier interface {
	ApplyGoods(state *model.GameState, goodsID model.MasterCardID, targets []int) error
}

type GoodsApplierService struct {
	GameMaster GameMasterService
}

func NewGoodsApplier(
	gameMaster GameMasterService,
) GoodsApplier {
	return &GoodsApplierService{
		GameMaster: gameMaster,
	}
}

func (s *GoodsApplierService) ApplyGoods(state *model.GameState, goodsID model.MasterCardID, targets []int) error {
	switch goodsID {
	case model.MasterCardID("oreilly"):
		return s.applyOreilly(state, targets)
	case model.MasterCardID("protein"):
		return s.applyProtein(state, targets)
	case model.MasterCardID("credit-card"):
		return s.applyCreditCard(state, targets)
	case model.MasterCardID("hackz-parker"):
		return s.applyHackzParker(state, targets)
	case model.MasterCardID("sake-bottle"):
		return s.applySakeBottle(state, targets)
	case model.MasterCardID("energy-drink"):
		return s.applyEnergyDrink(state, targets)
	case model.MasterCardID("starbucks"):
		return s.applyStarbucks(state, targets)
	case model.MasterCardID("gopher-doll"):
		return s.applyGopherDoll(state)
	case model.MasterCardID("hot-reload"):
		return s.applyHotReload(state)
	case model.MasterCardID("recruitment-agency"):
		return s.applyRecruitmentAgency(state)
	case model.MasterCardID("programming-school"):
		return s.applyProgrammingSchool(state)
	case model.MasterCardID("lan-cable"):
		return s.applyLanCable(state, targets)
	case model.MasterCardID("hhkb"):
		return s.applyHhkb(state)
	case model.MasterCardID("macbook"):
		return s.applyMacbook(state)
	default:
		return xerrors.Errorf("unknown goods id: %s", goodsID)
	}
}

// この番、自分の[知識]ラムモンが使うワザの、相手のバトルポケモンへのダメージを+40する
func (s *GoodsApplierService) applyOreilly(state *model.GameState, meta any) error {
	me := state.TurnPlayer

	me.Effect = append(me.Effect, &model.Effect{
		Trigger: "before-attack",
		Fn: func(any) (bool, error) {
			if me.BattleMonster.BaseCard.MasterCard.Type == model.MonsterTypeKnowledge {
				me.BattleMonster.SkillDamageAddition = 40
			}
			return true, nil
		},
	})

	me.Effect = append(me.Effect, &model.Effect{
		Trigger: "end-turn",
		Fn: func(any) (bool, error) {
			me.BattleMonster.SkillDamageAddition = 0
			return true, nil
		},
	})
	return nil
}

// 自分のエネルギーゾーンから[筋肉]エネルギーを2つ出し、自分の[筋肉]ラムモン1匹につける
func (s *GoodsApplierService) applyProtein(state *model.GameState, meta any) error {
	target := state.TurnPlayer.BattleMonster
	if !target.IsTypeEqual(model.MonsterTypeMuscle) {
		return xerrors.Errorf("target monster is not muscle type: %s", target.BaseCard.MasterCard.Type)
	}

	target.Energies = append(target.Energies, model.MonsterTypeMuscle, model.MonsterTypeMuscle)
	return nil
}

// 自分の[金]ラムモン1匹のHPを50回復
func (s *GoodsApplierService) applyCreditCard(state *model.GameState, meta any) error {
	target := state.TurnPlayer.BattleMonster
	if !target.IsTypeEqual(model.MonsterTypeMoney) {
		return xerrors.Errorf("target monster is not money type: %s", target.BaseCard.MasterCard.Type)
	}

	target.HP += 50
	return nil
}

// この番と次の相手の番、自分の[人気]ラムモン1匹は、ワザの追加効果や特性によるダメージを受けない。
func (s *GoodsApplierService) applyHackzParker(state *model.GameState, meta any) error {
	me := state.TurnPlayer

	target := me.BattleMonster
	if !target.IsTypeEqual(model.MonsterTypePopularity) {
		return xerrors.Errorf("target monster is not popularity type: %s", target.BaseCard.MasterCard.Type)
	}

	effectID := uuid.New().String()

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		ID:      effectID,
		Trigger: "take-effect",
		Fn: func(any) (bool, error) {
			return true, nil
		},
	})

	return nil
}

// 自分の[酒]ラムモンを1匹選ぶ。ウラが出るまでコインを投げ、オモテの数ぶんの[酒]エネルギーを自分のエネルギーゾーンから出し、そのラムモンにつける。
func (s *GoodsApplierService) applySakeBottle(state *model.GameState, meta any) error {
	target := state.TurnPlayer.BattleMonster
	if !target.IsTypeEqual(model.MonsterTypeAlchohol) {
		return xerrors.Errorf("target monster is not alchohol type: %s", target.BaseCard.MasterCard.Type)
	}

	var energies []model.MonsterType
	for s.GameMaster.FlipCoin() {
		energies = append(energies, model.MonsterTypeAlchohol)
	}
	target.Energies = append(target.Energies, energies...)

	return nil
}

// 自分のラムモン1匹のHPをすべて回復し、この番、そのラムモンが使うワザの、相手のバトルポケモンへのダメージを+30する。次の自分の番の開始時、そのラムモンのHPは0になる。
func (s *GoodsApplierService) applyEnergyDrink(state *model.GameState, meta any) error {
	me := state.TurnPlayer

	// HPを回復
	target := me.BattleMonster
	target.HP = target.BaseCard.MasterCard.HP

	// ダメージを+30
	target.SkillDamageAddition = 30

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "end-turn",
		Fn: func(any) (bool, error) {
			target.SkillDamageAddition = 0
			return true, nil
		},
	})

	// 次の自分の番の開始時、そのラムモンのHPは0になる
	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "start-my-turn",
		Fn: func(any) (bool, error) {
			if me.BattleMonster.ID == target.ID {
				target.HP = 0
				target.Knocked = true
				return true, nil
			}

			target, isFound := lo.Find(me.BenchMonsters, func(m *model.Monster) bool {
				return m.ID == target.ID
			})
			if !isFound {
				return true, nil
			}

			target.HP = 0

			return true, nil
		},
		UserID: me.UserID,
	})

	return nil
}

// 自分のラムモン1匹のHPを20回復
func (s *GoodsApplierService) applyStarbucks(state *model.GameState, meta any) error {
	target := state.TurnPlayer.BattleMonster
	target.HP += 20

	return nil
}

// この番、自分のバトルラムモンのにげるためのエネルギーを、1個少なくする。
func (s *GoodsApplierService) applyGopherDoll(state *model.GameState) error {
	me := state.TurnPlayer
	me.BattleMonster.RetreatCostAddition = -1

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "end-turn",
		Fn: func(any) (bool, error) {
			me.BattleMonster.RetreatCostAddition = 0
			return true, nil
		},
	})

	return nil
}

// 自分の手札をすべて山札に戻し、山札から同じ枚数のカードを引く。
func (s *GoodsApplierService) applyHotReload(state *model.GameState) error {
	count := len(state.TurnPlayer.Hands)

	state.TurnPlayer.Deck = append(state.TurnPlayer.Deck, state.TurnPlayer.Hands...)
	state.TurnPlayer.Hands = []*model.Card{}
	s.GameMaster.ShuffleDeck(state.TurnPlayer.Deck)

	s.GameMaster.DrawCards(state.TurnPlayer, count)

	return nil
}

// 自分の山札からたねラムモン以外のラムモンをランダムに1枚、手札に加える。
func (s *GoodsApplierService) applyRecruitmentAgency(state *model.GameState) error {
	for idx, card := range state.TurnPlayer.Deck {
		if card.MasterCard.CardType != model.CardTypeMonster || card.MasterCard.SubType == model.MonsterSubTypeBasic {
			continue
		}

		state.TurnPlayer.Hands = append(state.TurnPlayer.Hands, card)
		state.TurnPlayer.Deck = append(state.TurnPlayer.Deck[:idx], state.TurnPlayer.Deck[idx+1:]...)
		break
	}

	// 山札を見たためシャッフルする
	s.GameMaster.ShuffleDeck(state.TurnPlayer.Deck)

	return nil
}

// 自分の山札からたねラムモンをランダムに1枚、手札に加える。
func (s *GoodsApplierService) applyProgrammingSchool(state *model.GameState) error {
	for idx, card := range state.TurnPlayer.Deck {
		if card.MasterCard.CardType != model.CardTypeMonster || card.MasterCard.SubType != model.MonsterSubTypeBasic {
			continue
		}

		state.TurnPlayer.Hands = append(state.TurnPlayer.Hands, card)
		state.TurnPlayer.Deck = append(state.TurnPlayer.Deck[:idx], state.TurnPlayer.Deck[idx+1:]...)
		break
	}

	// 山札を見たためシャッフルする
	s.GameMaster.ShuffleDeck(state.TurnPlayer.Deck)

	return nil
}

// 自分のラムモン2匹を選び、そのラムモンについているエネルギーをすべて入れ替える。
func (s *GoodsApplierService) applyLanCable(state *model.GameState, meta any) error {
	var target1 *model.Monster
	var target2 *model.Monster

	target1.Energies, target2.Energies = target2.Energies, target1.Energies

	return nil
}

// 自分の山札から「駆け出しエンジニア」の進化先のラムモンをランダムに1枚、手札に加える。
func (s *GoodsApplierService) applyHhkb(state *model.GameState) error {
	for idx, card := range state.TurnPlayer.Deck {
		if card.MasterCard.CardType != model.CardTypeMonster ||
			!slices.Contains(card.MasterCard.EvolvesFrom, model.MasterCardID("beginner-engineer")) {
			continue
		}

		state.TurnPlayer.Hands = append(state.TurnPlayer.Hands, card)
		state.TurnPlayer.Deck = slices.Delete(state.TurnPlayer.Deck, idx, idx+1)
		break
	}
	return nil
}

// 「駆け出しエンジニア」とその進化先のラムモンが使うワザの、相手のバトルポケモンへのダメージを+20する
func (s *GoodsApplierService) applyMacbook(state *model.GameState) error {
	me := state.TurnPlayer

	me.Effect = append(me.Effect, &model.Effect{
		Trigger: "before-attack",
		Fn: func(any) (bool, error) {
			if me.BattleMonster.BaseCard.MasterCard.MasterCardID == model.MasterCardID("beginner-engineer") ||
				slices.Contains(me.BattleMonster.BaseCard.MasterCard.EvolvesFrom, model.MasterCardID("beginner-engineer")) {
				me.BattleMonster.SkillDamageAddition = 20
			}

			return true, nil
		},
	})

	me.Effect = append(me.Effect, &model.Effect{
		Trigger: "end-turn",
		Fn: func(any) (bool, error) {
			me.BattleMonster.SkillDamageAddition = 0
			return true, nil
		},
	})

	return nil
}
