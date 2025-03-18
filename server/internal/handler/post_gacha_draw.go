package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type PostGachaDrawHandler struct {
	drawGachaUsecase usecase.DrawGachaInputPort
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
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cards)
}
