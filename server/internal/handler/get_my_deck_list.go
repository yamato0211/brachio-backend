package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetDeckListHandler struct {
	getMyDecksUsecase usecase.GetMyDecksInputPort
}

func NewGetDeckListHandler(getMyDecksUsecase usecase.GetMyDecksInputPort) GetDeckListHandler {
	return GetDeckListHandler{
		getMyDecksUsecase: getMyDecksUsecase,
	}
}

func (h *GetDeckListHandler) GetDeckList(c echo.Context) error {
	userID := middleware.GetUserID(c)
	decks, err := h.getMyDecksUsecase.Execute(c.Request().Context(), userID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := make([]*schema.DeckBaseWithId, 0, len(decks))
	for _, deck := range decks {
		tc, err := schema.FactoryCard(*deck.ThumbnailCard)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		resp = append(resp, &schema.DeckBaseWithId{
			Id:            lo.ToPtr(deck.DeckID.String()),
			Name:          deck.Name,
			Color:         schema.Element(deck.Color),
			ThumbnailCard: *tc,
		})
	}
	return c.JSON(http.StatusOK, resp)
}
