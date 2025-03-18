package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type PostDeckHandler struct {
	createMyDeckUsecase usecase.CreateMyDeckInputPort
}

func (h *PostDeckHandler) CreateNewDeck(c echo.Context) error {
	userID := middleware.GetUserID(c)
	deck, err := h.createMyDeckUsecase.Execute(c.Request().Context(), userID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := &schema.CreateNewDeck201Response{
		Id: deck.DeckID.String(),
	}
	return c.JSON(http.StatusCreated, resp)
}
