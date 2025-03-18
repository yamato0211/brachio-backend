package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type PostGachaDrawHandler struct {
	drawGachaUsecase usecase.DrawGachaInputPort
}

func NewPostGachaDrawHandler(drawGachaUsecase usecase.DrawGachaInputPort) PostGachaDrawHandler {
	return PostGachaDrawHandler{
		drawGachaUsecase: drawGachaUsecase,
	}
}

func (h *PostGachaDrawHandler) DrawGacha(c echo.Context, gachaId string) error {
	var payload schema.DrawGachaRequest
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cards, err := h.drawGachaUsecase.Execute(c.Request().Context(), &usecase.DrawGachaInput{
		IsTen: payload.IsTenDraw,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cards)
}
