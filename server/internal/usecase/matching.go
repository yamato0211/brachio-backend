package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
)

type MatchingInputPort interface {
	Execute(ctx context.Context, input *MatchingInput) (roomID string, err error)
}

type MatchingInput struct {
	UserID   string
	DeckID   string
	Password string
}

type MatchingInteractor struct {
	GameStateRepository  repository.GameStateRepository
	DeckRepository       repository.DeckRepository
	MasterCardRepository repository.MasterCardRepository
	Matcher              service.Matcher
	GameMaster           service.GameMasterService
}

func NewMatchingUsecase(
	gsr repository.GameStateRepository,
	dr repository.DeckRepository,
	mcr repository.MasterCardRepository,
	m service.Matcher,
	gm service.GameMasterService,
) MatchingInputPort {
	return &MatchingInteractor{
		GameStateRepository:  gsr,
		DeckRepository:       dr,
		MasterCardRepository: mcr,
		Matcher:              m,
		GameMaster:           gm,
	}
}

func (i *MatchingInteractor) Execute(ctx context.Context, input *MatchingInput) (string, error) {
	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return "", err
	}

	fmt.Println("userID: ", userID)

	deckID, err := model.ParseDeckID(input.DeckID)
	if err != nil {
		return "", err
	}

	fmt.Println("deckID: ", deckID)

	// ユーザーのデッキを取得する
	deck, err := i.DeckRepository.Find(ctx, deckID)
	if err != nil {
		return "", err
	}

	masterCards, err := i.MasterCardRepository.FindAll(ctx)
	if err != nil {
		return "", err
	}

	myCards := make([]*model.MasterCard, 0, len(deck.MasterCardIDs))
	for _, cid := range deck.MasterCardIDs {
		mc, ok := lo.Find(masterCards, func(item *model.MasterCard) bool { return item.MasterCardID == cid })
		if !ok {
			return "", errors.New("master card not found")
		}
		myCards = append(myCards, mc)
	}
	deck.MasterCards = myCards

	fmt.Printf("deck: %+v\n", deck)

	wg := sync.WaitGroup{}
	wg.Add(1)

	var roomID model.RoomID
	err = i.Matcher.Apply(ctx, input.Password, func(_roomID model.RoomID) {
		defer wg.Done()

		roomID = _roomID

		var both bool
		err := i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
			state, err := i.GameStateRepository.Find(ctx, roomID)
			if err != nil && !errors.Is(err, model.ErrRoomNotFound) {
				return err
			}
			fmt.Printf("deck: %+v\n", deck)

			if state == nil {
				state = &model.GameState{
					RoomID: roomID,
					TurnPlayer: &model.Player{
						UserID:   userID,
						BaseDeck: deck,
					},
				}

				if err := i.GameStateRepository.Store(ctx, state); err != nil {
					return err
				}
			}

			state.NonTurnPlayer = &model.Player{
				UserID:   userID,
				BaseDeck: deck,
			}
			both = true

			if err := i.GameStateRepository.Store(ctx, state); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			log.Printf("transaction error: %v", err)
		}

		if both {
			if err := i.GameMaster.Matched(ctx, roomID); err != nil {
				log.Printf("game master matched error: %v", err)
				return
			}
		}

	})
	if err != nil {
		return "", err
	}

	wg.Wait()

	return roomID.String(), nil
}
