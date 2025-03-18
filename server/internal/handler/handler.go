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
	// DI
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	dc, err := dynamo.New(context.Background(), cfg.Dynamo)
	if err != nil {
		log.Fatal(err)
	}
	masterCardRepo := db.NewMasterCardRepository(*dc)
	drawGachaUsecase := usecase.NewDrawGachaUsecase(masterCardRepo)
	postGachaDrawHandler := NewPostGachaDrawHandler(drawGachaUsecase)
	return &Handler{
		GetMyCardListHandler:  GetMyCardListHandler{},
		GetDeckListHandler:    GetDeckListHandler{},
		PostDeckHandler:       PostDeckHandler{},
		GetDeckHandler:        GetDeckHandler{},
		PutDeckHandler:        PutDeckHandler{},
		DeleteDeckHandler:     DeleteDeckHandler{},
		GetMyItemListHandler:  GetMyItemListHandler{},
		GetGachaPowerHandler:  GetGachaPowerHandler{},
		GetGachaListHandler:   GetGachaListHandler{},
		PostGachaDrawHandler:  postGachaDrawHandler,
		GetWebSocketHandler:   GetWebSocketHandler{},
		GetHealthCheckHandler: GetHealthCheckHandler{},
	}
}
