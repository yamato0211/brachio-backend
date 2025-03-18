package usecase

import (
	"context"
	"errors"
	"log"

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
	GameStateRepository repository.GameStateRepository
	DeckRepository      repository.DeckRepository
	Matcher             service.Matcher
	GameMaster          service.GameMasterService
}

func NewMatchingUsecase(
	gsr repository.GameStateRepository,
	dr repository.DeckRepository,
	m service.Matcher,
	gm service.GameMasterService,
) MatchingInputPort {
	return &MatchingInteractor{
		GameStateRepository: gsr,
		DeckRepository:      dr,
		Matcher:             m,
		GameMaster:          gm,
	}
}

func (i *MatchingInteractor) Execute(ctx context.Context, input *MatchingInput) (string, error) {
	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return "", err
	}

	deckID, err := model.ParseDeckID(input.DeckID)
	if err != nil {
		return "", err
	}

	// ユーザーのデッキを取得する
	deck, err := i.DeckRepository.Find(ctx, deckID)
	if err != nil {
		return "", err
	}

	ch := make(chan string)

	err = i.Matcher.Apply(ctx, input.Password, func(roomID model.RoomID) {
		var both bool
		err := i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
			state, err := i.GameStateRepository.Find(ctx, roomID)
			if err != nil && !errors.Is(err, model.ErrRoomNotFound) {
				return err
			}

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

		ch <- roomID.String()
	})
	if err != nil {
		return "", err
	}

	return <-ch, nil
}
