package service

import (
	"fmt"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"golang.org/x/xerrors"
)

type SkillApprier interface {
	ApplySkill(p1, p2 *model.Player, id int) (int, error)
}

type SkillApprierService struct{}

func NewSkillApprier() SkillApprier {
	return &SkillApprierService{}
}

func (s *SkillApprierService) ApplySkill(p1 *model.Player, p2 *model.Player, id int) (int, error) {
	attackMonster := p1.BattleMonster

	if len(attackMonster.BaseCard.MasterCard.Skills) <= id {
		return 0, xerrors.Errorf("invalid attack id: %d", id)
	}

	skill := attackMonster.BaseCard.MasterCard.Skills[id]
	skillID := fmt.Sprintf("%s-%d", attackMonster.BaseCard.MasterCard.MasterCardID, id+1)

	// TODO: Cost Check

	switch skillID {
	case "kizuku-1":
		return s.applyKizuku(p1, p2, skill)
	default:
		return s.applyDefault(p1, p2, skill)
	}

}

func (s *SkillApprierService) applyDefault(p1, p2 *model.Player, skill *model.Skill) (int, error) {
	m1 := p1.BattleMonster
	m2 := p2.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction

	m2.HP -= damage
	m2.HP = max(0, m2.HP)
	if m2.HP == 0 {
		m2.Knocked = true
	}

	return damage, nil
}

// コインを1回投げ表なら相手のベンチラムモン全員にも200ダメージ、裏ならこのラムモンについているエネルギーをすべてトラッシュする
func (s *SkillApprierService) applyKizuku(p1, p2 *model.Player, skill *model.Skill) (int, error) {
	m1 := p1.BattleMonster
	m2 := p2.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction

	p1.Effect = append(p1.Effect, &model.Effect{
		Trigger: "after-coin",
		Fn: func(x any) (bool, error) {
			m2.HP -= damage
			m2.HP = max(0, m2.HP)
			if m2.HP == 0 {
				m2.Knocked = true
			}

			coin, ok := x.(bool)
			if !ok {
				return true, xerrors.Errorf("invalid coin: %v", x)
			}

			if coin {
				for _, monster := range p2.BenchMonsters {
					monster.HP -= 200
					monster.HP = max(0, monster.HP)
					if monster.HP == 0 {
						monster.Knocked = true
					}
				}
			} else {
				m1.Energies = make([]model.MonsterType, 0)
			}

			return true, nil
		},
	})

	return damage, nil
}
