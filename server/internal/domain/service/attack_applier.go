package service

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"golang.org/x/xerrors"
)

type SkillApprier interface {
	ApplySkill(state *model.GameState, id int) (int, error)
}

type SkillApprierService struct{}

func NewSkillApprier() SkillApprier {
	return &SkillApprierService{}
}

func (s *SkillApprierService) ApplySkill(state *model.GameState, id int) (int, error) {
	attackMonster := state.TurnPlayer.BattleMonster

	if len(attackMonster.BaseCard.MasterCard.Skills) <= id {
		return 0, xerrors.Errorf("invalid attack id: %d", id)
	}

	skill := attackMonster.BaseCard.MasterCard.Skills[id]
	skillID := fmt.Sprintf("%s-%d", attackMonster.BaseCard.MasterCard.MasterCardID, id+1)

	cost := lo.GroupBy(skill.Cost, func(e model.MonsterType) model.MonsterType { return e })
	for k, v := range cost {
		if lo.Count(attackMonster.Energies, k) < len(v) {
			return 0, xerrors.Errorf("not enough energy: %v", k)
		}
	}

	switch skillID {
	case "kizuku-1":
		return s.applyKizuku(state, skill)
	default:
		return s.applyDefault(state, skill)
	}

}

func (s *SkillApprierService) applyDefault(state *model.GameState, skill *model.Skill) (int, error) {
	m1 := state.TurnPlayer.BattleMonster
	m2 := state.NonTurnPlayer.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction

	m2.HP -= damage
	m2.HP = max(0, m2.HP)
	if m2.HP == 0 {
		m2.Knocked = true
	}

	return damage, nil
}

// コインを1回投げ表なら相手のベンチラムモン全員にも200ダメージ、裏ならこのラムモンについているエネルギーをすべてトラッシュする
func (s *SkillApprierService) applyKizuku(state *model.GameState, skill *model.Skill) (int, error) {
	m1 := state.TurnPlayer.BattleMonster
	m2 := state.NonTurnPlayer.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction
	state.Damages = make([]int, 4)
	state.Damages[0] = damage

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "after-coin",
		Fn: func(state *model.GameState, x any) (bool, error) {
			state.NonTurnPlayer.BattleMonster.HP -= damage
			state.NonTurnPlayer.BattleMonster.HP = max(0, state.NonTurnPlayer.BattleMonster.HP)
			if state.NonTurnPlayer.BattleMonster.HP == 0 {
				state.NonTurnPlayer.BattleMonster.Knocked = true
			}

			coin, ok := x.(bool)
			if !ok {
				return true, xerrors.Errorf("invalid coin: %v", x)
			}

			if coin {
				for i, monster := range state.NonTurnPlayer.BenchMonsters {
					monster.HP -= 200
					monster.HP = max(0, monster.HP)
					if monster.HP == 0 {
						monster.Knocked = true
					}

					state.Damages[i+1] = 200
				}
			} else {
				state.TurnPlayer.BattleMonster.Energies = make([]model.MonsterType, 0)
			}

			return true, nil
		},
	})

	return damage, nil
}
