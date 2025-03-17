package handler

import "github.com/labstack/echo/v4"

type DeleteDeckHandler struct{}

func (h *DeleteDeckHandler) DeleteDeck(c echo.Context, deckId string) error {
	return nil
}
