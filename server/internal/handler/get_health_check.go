package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetHealthCheckHandler struct{}

func (h *GetHealthCheckHandler) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
