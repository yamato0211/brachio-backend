package service

import (
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"golang.org/x/xerrors"
)

type SupporterApplier interface {
	ApplySupporter(state *model.GameState, goodsID model.MasterCardID, targets []int) error
}

type SupporterApplierService struct {
	GameMaster GameMasterService
}

func NewSupporterApplier(
	gameMaster GameMasterService,
) SupporterApplier {
	return &SupporterApplierService{
		GameMaster: gameMaster,
	}
}

func (s *SupporterApplierService) ApplySupporter(state *model.GameState, goodsID model.MasterCardID, targets []int) error {
	switch goodsID {
	case model.MasterCardID("chat-gpt"):
		return s.applyChatGpt(state, targets)
	case model.MasterCardID("spaghetti-code"):
		return s.applySpaghettiCode(state, targets)
	case model.MasterCardID("flaming-project"):
		return s.applyFlamingProject(state, targets)
	case model.MasterCardID("security-soft"):
		return s.applySecuritySoft(state, targets)
	case model.MasterCardID("strict-mode"):
		return s.applyStrictMode(state, targets)
	case model.MasterCardID("firewall"):
		return s.applyFirewall(state, targets)
	default:
		return xerrors.Errorf("unknown supporter id: %s", goodsID)
	}
}

// この番、自分のバトルラムモンのにげるためのエネルギーを、2個少なくする。
func (s *SupporterApplierService) applyChatGpt(state *model.GameState, meta any) error {
	me := state.TurnPlayer

	me.BattleMonster.RetreatCostAddition -= 2

	me.Effect = append(me.Effect, &model.Effect{
		Trigger: "turn-end",
		Fn: func(any) (bool, error) {
			me.BattleMonster.RetreatCostAddition = 0
			return true, nil
		},
	})

	return nil
}

// 相手のバトルラムモンのランダムなエネルギー1個を、ランダムなエネルギーに変える。
func (s *SupporterApplierService) applySpaghettiCode(state *model.GameState, meta any) error {
	return nil
}

// お互いのバトルラムモンについているエネルギーをすべてトラッシュする。
func (s *SupporterApplierService) applyFlamingProject(state *model.GameState, meta any) error {
	return nil
}

// 相手の手札からランダムに1枚トラッシュ
func (s *SupporterApplierService) applySecuritySoft(state *model.GameState, meta any) error {
	return nil
}

// 自分の山札を2枚引く。
func (s *SupporterApplierService) applyStrictMode(state *model.GameState, meta any) error {
	return nil
}

// 次の相手の番、自分のラムモン全員が、相手のラムモンから受けるダメージを-20する。
func (s *SupporterApplierService) applyFirewall(state *model.GameState, meta any) error {
	return nil
}
