package handler

import "github.com/labstack/echo/v4"

type GetMyDeckHandler struct{}

func (h *GetMyDeckHandler) GetMyDeck(c echo.Context, deckId string) error {
	return nil
}
