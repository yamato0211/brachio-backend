package service

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	"golang.org/x/xerrors"
)

type SkillApprier interface {
	ApplySkill(ctx context.Context, state *model.GameState, id int) (int, error)
}

type SkillApprierService struct {
	GameMaster      GameMasterService
	GameEventSender GameEventSender
}

func NewSkillApprier(
	gm GameMasterService,
	ges GameEventSender,
) SkillApprier {
	return &SkillApprierService{
		GameMaster:      gm,
		GameEventSender: ges,
	}
}

func (s *SkillApprierService) ApplySkill(ctx context.Context, state *model.GameState, id int) (int, error) {
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
		return s.applyKizuku1(ctx, state, skill)
	case "dolly-2":
		return s.applyDolly2(ctx, state, skill)
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
func (s *SkillApprierService) applyKizuku1(ctx context.Context, state *model.GameState, skill *model.Skill) (int, error) {
	m1 := state.TurnPlayer.BattleMonster
	m2 := state.NonTurnPlayer.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "after-coin",
		Fn: func(state *model.GameState, x any) (bool, error) {
			state.NonTurnPlayer.BattleMonster.HP -= damage
			state.NonTurnPlayer.BattleMonster.HP = max(0, state.NonTurnPlayer.BattleMonster.HP)
			if state.NonTurnPlayer.BattleMonster.HP == 0 {
				state.NonTurnPlayer.BattleMonster.Knocked = true
			}
			state.Damages = make([]int, 4)
			state.Damages[0] = damage

			coin, ok := x.([]bool)
			if !ok {
				return true, xerrors.Errorf("invalid coin: %v", x)
			}

			if coin[0] {
				for i, monster := range state.NonTurnPlayer.BenchMonsters {
					monster.HP -= 200
					monster.HP = max(0, monster.HP)
					if monster.HP == 0 {
						monster.Knocked = true
					}

					state.Damages[i+1] = 200
				}

				damage := lo.Map(state.Damages, func(d int, _ int) int32 { return int32(d) })
				eventForMe := &messages.EffectWithSecret{Effect: &messages.EffectWithSecret_Damage{Damage: &messages.DamageEffect{Position: 0, Amount: damage}}}
				if err := s.GameEventSender.SendDrawEffectEventToActor(ctx, state.TurnPlayer.UserID, eventForMe); err != nil {
					return true, err
				}
				eventForOppo := &messages.Effect{Effect: &messages.Effect_Damage{Damage: &messages.DamageEffect{Position: 0, Amount: damage}}}
				if err := s.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, eventForOppo); err != nil {
					return true, err
				}

			} else {
				state.TurnPlayer.BattleMonster.Energies = make([]model.MonsterType, 0)

				eventForMe := &messages.EffectWithSecret{Effect: &messages.EffectWithSecret_EnergyTrash{EnergyTrash: &messages.EnergyTrashEffect{Position: 0, Energy: []messages.Element{}}}}
				if err := s.GameEventSender.SendDrawEffectEventToActor(ctx, state.TurnPlayer.UserID, eventForMe); err != nil {
					return true, err
				}
				eventForOppo := &messages.Effect{Effect: &messages.Effect_EnergyTrash{EnergyTrash: &messages.EnergyTrashEffect{Position: 0, Energy: []messages.Element{}}}}
				if err := s.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, eventForOppo); err != nil {
					return true, err
				}
			}

			return true, nil
		},
	})

	return damage, nil
}

// コインを2回投げ2回ともウラなら、このポケモンにも100ダメージ
func (s *SkillApprierService) applyDolly2(ctx context.Context, state *model.GameState, skill *model.Skill) (int, error) {
	m1 := state.TurnPlayer.BattleMonster
	m2 := state.NonTurnPlayer.BattleMonster

	damage := skill.Damage + m1.SkillDamageAddition - m2.SkillDamageReduction

	state.TurnPlayer.Effect = append(state.TurnPlayer.Effect, &model.Effect{
		Trigger: "after-coin",
		Fn: func(state *model.GameState, x any) (bool, error) {
			state.NonTurnPlayer.BattleMonster.HP -= damage
			state.NonTurnPlayer.BattleMonster.HP = max(0, state.NonTurnPlayer.BattleMonster.HP)
			if state.NonTurnPlayer.BattleMonster.HP == 0 {
				state.NonTurnPlayer.BattleMonster.Knocked = true
			}
			state.Damages = make([]int, 4)
			state.Damages[0] = damage

			coin, ok := x.([]bool)
			if !ok {
				return true, xerrors.Errorf("invalid coin: %v", x)
			}

			if !coin[0] && !coin[1] {
				state.TurnPlayer.BattleMonster.HP -= 100
				state.NonTurnPlayer.BattleMonster.HP = max(0, state.NonTurnPlayer.BattleMonster.HP)
				if state.NonTurnPlayer.BattleMonster.HP == 0 {
					state.NonTurnPlayer.BattleMonster.Knocked = true
				}

				isGameSet, winner := s.GameMaster.CheckWin(state)
				if isGameSet {
					winOrLose := &messages.DecideWinOrLoseEffect{UserId: winner.String()}
					if err := s.GameEventSender.SendDrawEffectEventToActor(ctx, state.TurnPlayer.UserID, &messages.EffectWithSecret{Effect: &messages.EffectWithSecret_DecideWinOrLose{DecideWinOrLose: winOrLose}}); err != nil {
						return true, err
					}
					if err := s.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, &messages.Effect{Effect: &messages.Effect_DecideWinOrLose{DecideWinOrLose: winOrLose}}); err != nil {
						return true, err
					}
				}
			}

			return true, nil
		},
	})

	return damage, nil
}
