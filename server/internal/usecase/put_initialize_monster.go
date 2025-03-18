package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"golang.org/x/xerrors"
)

type PutInitializeMonsterInputPort interface {
	Execute(ctx context.Context, input *PutInitializeMonsterInput) error
}

type PutInitializeMonsterInput struct {
	RoomID   string
	UserID   string
	CardID   string
	Position int
}

type PutInitializeMonsterInteractor struct {
	GameStateRepository repository.GameStateRepository
	GameManeger         service.GameMasterService
}

func NewPutInitializeMonsterUsecase(
	gsr repository.GameStateRepository,
	gm service.GameMasterService,
) PutInitializeMonsterInputPort {
	return &PutInitializeMonsterInteractor{
		GameStateRepository: gsr,
		GameManeger:         gm,
	}
}

func (i *PutInitializeMonsterInteractor) Execute(ctx context.Context, input *PutInitializeMonsterInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}
	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}
	cardID, err := model.ParseCardID(input.CardID)
	if err != nil {
		return err
	}

	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		me, err := state.FindMeByUserID(userID)
		if err != nil {
			return err
		}

		// カードが手札にあるか確認
		card, idx, isFound := lo.FindIndexOf(me.Hands, func(c *model.Card) bool {
			return c.MasterCard.SubType == model.MonsterSubTypeBasic && c.CardID == cardID
		})
		if !isFound {
			return xerrors.Errorf("card not found in hand: %d", cardID)
		}

		monster, err := card.Summon(1)
		if err != nil {
			return err
		}

		// Position 0 はバトルゾーン
		if input.Position == 0 {
			// バトルゾーンにカードを出す
			if me.BattleMonster != nil {
				return xerrors.Errorf("battle monster already exists")
			}
			me.BattleMonster = monster
		} else {
			// 控えに出す
			if me.BenchMonsters[input.Position-1] != nil {
				return xerrors.Errorf("monster already exists in bench: %d", input.Position)
			}
			me.BattleMonster = monster
		}

		// 手札からカードを削除
		me.Hands = slices.Delete(me.Hands, idx, idx+1)

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	return nil
}
