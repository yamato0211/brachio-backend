package handler

import (
	"context"
	"log"

	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/gateway/db"
	"github.com/yamato0211/brachio-backend/internal/infra/dynamo"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type Handler struct {
	// Card
	GetMyCardListHandler

	// Deck
	GetDeckListHandler
	PostDeckHandler
	GetDeckHandler
	PutDeckHandler
	DeleteDeckHandler

	// Item
	GetMyItemListHandler

	// Gacha
	GetGachaPowerHandler
	GetGachaListHandler
	PostGachaDrawHandler

	// Game
	GetWebSocketHandler

	// HealthCheck
	GetHealthCheckHandler
}

func New() *Handler {
	ctx := context.Background()
	// DI
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	dc, err := dynamo.New(ctx, cfg.Dynamo)
	if err != nil {
		log.Fatal(err)
	}

	// repo
	masterCardRepo := db.NewMasterCardRepository(dc)
	deckRepo := db.NewDeckRepository(dc)
	masterItemRepo := db.NewMasterItemRepository(dc)
	userRepo := db.NewUserRepository(dc)

	// usecase
	drawGachaUsecase := usecase.NewDrawGachaUsecase(masterCardRepo)
	getMyDecksUsecase := usecase.NewGetMyDecksUsecase(deckRepo, masterCardRepo)
	getMyItemsUsecase := usecase.NewGetMyItemsUsecase(masterItemRepo, userRepo)
	createMyDeckUsecase := usecase.NewCreateMyDeckUsecase(deckRepo)
	getMyDeckUsecase := usecase.NewGetMyDeckUsecase(deckRepo, masterCardRepo)
	updateMyDeckUsecase := usecase.NewUpdateMyDeckUsecase(deckRepo)
	deleteMyDeckUsecase := usecase.NewDeleteMyDeckUsecase(deckRepo)
	getMyCardsUsecase := usecase.NewGetMyCardsUsecase(masterCardRepo, userRepo)

	return &Handler{
		GetMyCardListHandler: GetMyCardListHandler{
			getMyCardsUsecase: getMyCardsUsecase,
		},
		GetDeckListHandler: GetDeckListHandler{
			getMyDecksUsecase: getMyDecksUsecase,
		},
		PostDeckHandler: PostDeckHandler{
			createMyDeckUsecase: createMyDeckUsecase,
		},
		GetDeckHandler: GetDeckHandler{
			getMyDeckUsecase: getMyDeckUsecase,
		},
		PutDeckHandler: PutDeckHandler{
			updateMyDeckUsecase: updateMyDeckUsecase,
		},
		DeleteDeckHandler: DeleteDeckHandler{
			DeleteMyDeckUsecase: deleteMyDeckUsecase,
		},
		GetMyItemListHandler: GetMyItemListHandler{
			getMyItemsUsecase: getMyItemsUsecase,
		},
		GetGachaPowerHandler: GetGachaPowerHandler{},
		GetGachaListHandler:  GetGachaListHandler{},
		PostGachaDrawHandler: PostGachaDrawHandler{
			drawGachaUsecase: drawGachaUsecase,
		},
		GetWebSocketHandler:   GetWebSocketHandler{},
		GetHealthCheckHandler: GetHealthCheckHandler{},
	}
}
