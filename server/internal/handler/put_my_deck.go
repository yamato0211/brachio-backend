package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type PutDeckHandler struct {
	updateMyDeckUsecase usecase.UpdateMyDeckInputPort
}

func NewPutDeckHandler(updateMyDeckUsecase usecase.UpdateMyDeckInputPort) *PutDeckHandler {
	return &PutDeckHandler{updateMyDeckUsecase: updateMyDeckUsecase}
}

func (h *PutDeckHandler) UpdateDeck(c echo.Context, deckId string) error {
	var payload schema.UpdateDeck
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	dki, err := model.ParseDeckID(deckId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: 所有権の確認

	input := &model.Deck{
		DeckID: dki,
		Name:   payload.Name,
		Color:  model.MonsterType(payload.Color),
		Energies: lo.Map(payload.Energies, func(e schema.Element, _ int) model.MonsterType {
			return model.MonsterType(e)
		}),
		MasterCardIDs: lo.Map(payload.MasterCardIds, func(id string, _ int) model.MasterCardID {
			mci, _ := model.ParseMasterCardID(id)
			return mci
		}),
		ThumbnailCardID: lo.Must(model.ParseMasterCardID(payload.ThumbnailCardId)),
	}
	if err := h.updateMyDeckUsecase.Execute(c.Request().Context(), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
