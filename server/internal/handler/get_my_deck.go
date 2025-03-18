package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/guregu/dynamo/v2"
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetDeckHandler struct {
	getMyDeckUsecase usecase.GetMyDeckInputPort
}

func (h *GetDeckHandler) GetDeck(c echo.Context, deckId string) error {
	deck, err := h.getMyDeckUsecase.Execute(c.Request().Context(), model.DeckID(deckId))
	if errors.Is(err, dynamo.ErrNotFound) {
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprintf("deck %s not found", deckId)})
	}

	resp, err := schema.DeckWithIdFromEntity(deck)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
