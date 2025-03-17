package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/infra/dynamo"
)

type GetMyCardListHandler struct{}

func (h *GetMyCardListHandler) GetCards(c echo.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	dc, err := dynamo.New(context.Background(), cfg.Dynamo)
	if err != nil {
		slog.Warn("failed to create dynamo client", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
	}

	type User struct {
		ID   string `dynamo:"UserId,hash"`
		Name string `dynamo:"Name"`
	}
	users := []User{}
	err = dc.Table("Users").Scan().All(c.Request().Context(), &users)
	if err != nil {
		slog.Error("error!", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
