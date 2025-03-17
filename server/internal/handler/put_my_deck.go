package handler

import "github.com/labstack/echo/v4"

type PutMyDeckHandler struct{}

func (h *PutMyDeckHandler) PutMyDeck(c echo.Context, deckId string) error {
	return nil
}
