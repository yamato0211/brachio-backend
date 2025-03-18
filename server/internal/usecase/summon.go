package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"golang.org/x/xerrors"
)

type SummonInputPort interface {
	Execute(ctx context.Context, input *SummonInput) error
}

type SummonInput struct {
	RoomID   string
	UserID   string
	CardID   string
	Position int
}

type SummonInteractor struct {
	GameStateRepository repository.GameStateRepository
}

func NewSummonUsecase(
	gsr repository.GameStateRepository,
) SummonInputPort {
	return &SummonInteractor{
		GameStateRepository: gsr,
	}
}

func (i *SummonInteractor) Execute(ctx context.Context, input *SummonInput) error {
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

		me := state.FindPlayerByUserID(userID)
		if me == nil {
			return xerrors.Errorf("player not found: %s", userID)
		}

		card, idx, isFound := lo.FindIndexOf(me.Hands, func(c *model.Card) bool {
			return c.MasterCard.SubType == model.MonsterSubTypeBasic && c.CardID == cardID
		})
		if !isFound {
			return xerrors.Errorf("card not found: %d", cardID)
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
			me.BenchMonsters[input.Position-1] = monster
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
