package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
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
	userID := middleware.GetUserID(c)
	cards, err := h.drawGachaUsecase.Execute(c.Request().Context(), &usecase.DrawGachaInput{
		IsTen:  payload.IsTenDraw,
		UserID: userID,
	})
	if errors.Is(err, model.ErrNoEnoughPackPower) {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	mcs := make([]schema.MasterCard, 0, len(cards))
	for _, card := range cards {
		mc, err := schema.MasterCardWithFromEntity(card)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		mcs = append(mcs, *mc)
	}

	resp := make([]*schema.Pack, 0, 10)
	for _, pack := range lo.Chunk(mcs, 5) {
		resp = append(resp, &schema.Pack{
			Cards: lo.ToPtr(pack),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
