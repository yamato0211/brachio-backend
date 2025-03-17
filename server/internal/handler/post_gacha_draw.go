package handler

import "github.com/labstack/echo/v4"

type PostGachaDrawHandler struct{}

func (h *PostGachaDrawHandler) DrawGacha(c echo.Context, gachaId string) error {
	return nil
}
