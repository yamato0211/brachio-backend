package handler

import "github.com/labstack/echo/v4"

type GetMyCardHandler struct{}

func (h *GetMyCardHandler) GetMyCard(c echo.Context, cardNumber string) error {
	return nil
}
