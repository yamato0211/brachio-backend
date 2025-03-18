package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetUserHandler struct {
	getUserUsecase usecase.GetUserInputPort
}

func (h *GetUserHandler) GetUser(c echo.Context, userID string) error {
	user, err := h.getUserUsecase.Execute(c.Request().Context(), userID)
	if errors.Is(err, model.ErrUserNotFound) {
		return c.JSON(http.StatusNotFound, err)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, schema.User{
		ImageUrl: user.ImageURL,
		Name:     user.Name,
	})
}
