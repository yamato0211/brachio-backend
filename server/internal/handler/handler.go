package handler

type Handler struct {
	// Card
	GetMyCardListHandler
	GetMyCardHandler

	// Deck
	GetMyDeckListHandler
	PostMyDeckHandler
	GetMyDeckHandler
	PutMyDeckHandler

	// Item
	GetMyItemListHandler

	// Gacha
	GetPackPowerHandler
	GetGachaListHandler
	PostGachaDrawHandler
}

func New() *Handler {
	return &Handler{
		GetMyCardListHandler: GetMyCardListHandler{},
		GetMyCardHandler:     GetMyCardHandler{},
		GetMyDeckListHandler: GetMyDeckListHandler{},
		PostMyDeckHandler:    PostMyDeckHandler{},
		GetMyDeckHandler:     GetMyDeckHandler{},
		PutMyDeckHandler:     PutMyDeckHandler{},
		GetMyItemListHandler: GetMyItemListHandler{},
		GetPackPowerHandler:  GetPackPowerHandler{},
		GetGachaListHandler:  GetGachaListHandler{},
		PostGachaDrawHandler: PostGachaDrawHandler{},
	}
}
