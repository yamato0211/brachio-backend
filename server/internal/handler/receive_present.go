package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type ReceivePresentHandler struct {
	receivePresentUsecase usecase.ReceivePresentInputPort
}

func (h *ReceivePresentHandler) ReceivePresent(c echo.Context, presentID string) error {
	userID := middleware.GetUserID(c)
	err := h.receivePresentUsecase.Execute(c.Request().Context(), userID, presentID)
	if errors.Is(err, model.ErrAlreadyReceivedPresent) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
