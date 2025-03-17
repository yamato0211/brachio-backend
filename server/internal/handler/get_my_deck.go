package handler

import "github.com/labstack/echo/v4"

type GetDeckHandler struct{}

func (h *GetDeckHandler) GetDeck(c echo.Context, deckId string) error {
	return nil
}
