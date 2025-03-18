package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyCardsInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.MasterCardWithCount, error)
}

type GetMyCardsInteractor struct {
	masterCardRepository repository.MasterCardRepository
	userRepository       repository.UserRepository
}

func (g *GetMyCardsInteractor) Execute(ctx context.Context, userID string) ([]*model.MasterCardWithCount, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}

	user, err := g.userRepository.Find(ctx, uid)
	if err != nil {
		return nil, err
	}

	masterCards, err := g.masterCardRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var myCards []*model.MasterCardWithCount
	for _, card := range masterCards {
		myCards = append(myCards, &model.MasterCardWithCount{
			MasterCard: card,
			Count:      user.CardIDsWithCount[card.MasterCardID.String()],
		})
	}
	return myCards, nil
}

func NewGetMyCardsUsecase(mcr repository.MasterCardRepository, ur repository.UserRepository) GetMyCardsInputPort {
	return &GetMyCardsInteractor{
		masterCardRepository: mcr,
		userRepository:       ur,
	}
}
