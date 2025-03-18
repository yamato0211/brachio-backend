package service

import (
	"context"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"golang.org/x/xerrors"
)

type AbilityApplier interface {
	ApplyAbility(ctx context.Context, state *model.GameState, monster *model.Monster) error
}

type AbilityApplierImpl struct{}

func NewAbilityApplier() AbilityApplier {
	return &AbilityApplierImpl{}
}

func (s *AbilityApplierImpl) ApplyAbility(ctx context.Context, state *model.GameState, monster *model.Monster) error {
	switch monster.BaseCard.MasterCard.MasterCardID {
	case model.MasterCardID("kizuku"):
		return s.ApplyKizukuAvility(ctx, state, monster)
	default:
		return xerrors.Errorf("unknown ability")
	}
}

// 自分の番に1回使える。自分のエネルギーゾーンからランダムにエネルギーを5個出し、このラムモンにつける。
func (s *AbilityApplierImpl) ApplyKizukuAvility(ctx context.Context, state *model.GameState, monster *model.Monster) error {
	if monster.AbilityUsed {
		return nil
	}

	energies := lo.Samples([]model.MonsterType{
		model.MonsterTypeMoney,
		model.MonsterTypeAlchohol,
		model.MonsterTypeKnowledge,
		model.MonsterTypeMuscle,
		model.MonsterTypePopularity,
	}, 5)

	monster.Energies = append(monster.Energies, energies...)

	monster.AbilityUsed = true
	return nil

}
