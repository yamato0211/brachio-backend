package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetMyCardListHandler struct {
	getMyCardsUsecase usecase.GetMyCardsInputPort
}

func (h *GetMyCardListHandler) GetCards(c echo.Context, params schema.GetCardsParams) error {
	isAll := lo.Ternary(params.IsAll != nil && *params.IsAll == 1, true, false)

	userID := middleware.GetUserID(c)

	cards, err := h.getMyCardsUsecase.Execute(c.Request().Context(), &usecase.GetMyCardsInput{
		UserID: userID,
		IsAll:  isAll,
	})
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := make([]*schema.MasterCardWithCount, 0, len(cards))
	for _, card := range cards {
		sc, err := schema.MasterCardWithFromEntity(card.MasterCard)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		resp = append(resp, &schema.MasterCardWithCount{
			MasterCard: *sc,
			Count:      card.Count,
		})
	}
	return c.JSON(http.StatusOK, resp)
}
