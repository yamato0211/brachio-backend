package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type DeleteDeckHandler struct {
	DeleteMyDeckUsecase usecase.DeleteMyDeckInputPort
}

func (h *DeleteDeckHandler) DeleteDeck(c echo.Context, deckId string) error {
	err := h.DeleteMyDeckUsecase.Execute(c.Request().Context(), deckId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
