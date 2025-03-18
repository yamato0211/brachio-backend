package usecase

import (
	"context"
	"fmt"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyItemsInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.MasterItemWithCount, error)
}

type GetMyItemsInteractor struct {
	itemRepo repository.MasterItemRepository
	userRepo repository.UserRepository
}

func (g *GetMyItemsInteractor) Execute(ctx context.Context, userID string) ([]*model.MasterItemWithCount, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}
	user, err := g.userRepo.Find(ctx, uid)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.ItemIDsWithCount)

	masterItems, err := g.itemRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var myItems []*model.MasterItemWithCount
	for _, item := range masterItems {
		fmt.Println(item)
		myItems = append(myItems, &model.MasterItemWithCount{
			MasterItem: item,
			Count:      user.ItemIDsWithCount[item.ItemID.String()],
		})
	}
	return myItems, nil
}

func NewGetMyItemsUsecase(ir repository.MasterItemRepository, ur repository.UserRepository) GetMyItemsInputPort {
	return &GetMyItemsInteractor{
		itemRepo: ir,
		userRepo: ur,
	}
}
