package usecase

import (
	"context"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyDecksInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.Deck, error)
}

type GetMyDecksInteractor struct {
	DeckRepository     repository.DeckRepository
	MasterCardRepitory repository.MasterCardRepository
}

func NewGetMyDecksUsecase(dr repository.DeckRepository, mcr repository.MasterCardRepository) GetMyDecksInputPort {
	return &GetMyDecksInteractor{
		DeckRepository:     dr,
		MasterCardRepitory: mcr,
	}
}

func (g *GetMyDecksInteractor) Execute(ctx context.Context, userID string) ([]*model.Deck, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}

	// 自分の持つデッキ
	decks, err := g.DeckRepository.FindAll(ctx, uid)
	if err != nil {
		return nil, err
	}

	// テンプレートデッキ
	templateDecks, err := g.DeckRepository.FindAllTempalte(ctx)
	if err != nil {
		return nil, err
	}

	masterCards, err := g.MasterCardRepitory.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	concatDecks := append(decks, templateDecks...)

	var resp = make([]*model.Deck, 0, len(concatDecks))
	for _, deck := range concatDecks {
		thumbnailCard, _ := lo.Find(masterCards, func(item *model.MasterCard) bool {
			return item.MasterCardID == deck.ThumbnailCardID
		})
		deck.ThumbnailCard = thumbnailCard
		resp = append(resp, deck)
	}
	return resp, nil
}
