package handler

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
		PostGachaDrawHandler:  PostGachaDrawHandler{},
		GetWebSocketHandler:   GetWebSocketHandler{},
		GetHealthCheckHandler: GetHealthCheckHandler{},
	}
}
