package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyPresentsInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.Present, error)
}

type GetMyPresentsInteractor struct {
	presentRepository repository.PresentRepository
	itemRepository    repository.MasterItemRepository
}

func NewGetMyPresentsUsecase(pr repository.PresentRepository, ir repository.MasterItemRepository) GetMyPresentsInputPort {
	return &GetMyPresentsInteractor{
		presentRepository: pr,
		itemRepository:    ir,
	}
}

func (g *GetMyPresentsInteractor) Execute(ctx context.Context, userID string) ([]*model.Present, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}

	presents, err := g.presentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	masterItems, err := g.itemRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(presents, func(i *model.Present, j *model.Present) int {
		return i.Time - j.Time
	})

	resp := make([]*model.Present, 0, len(presents))
	for _, p := range presents {
		// 未受け取りのプレゼントのみを取得
		if !slices.Contains(p.ReceivedUserIDs, uid) {
			item, ok := lo.Find(masterItems, func(i *model.MasterItem) bool {
				return i.ItemID == p.Item.ItemID
			})
			if !ok {
				return nil, model.ErrItemNotFound
			}
			p.Item = *item
			resp = append(resp, p)
		}
	}
	return resp, nil
}
