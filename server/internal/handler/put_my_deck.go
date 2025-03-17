package handler

import "github.com/labstack/echo/v4"

type PutDeckHandler struct{}

func (h *PutDeckHandler) UpdateDeck(c echo.Context, deckId string) error {
	return nil
}
