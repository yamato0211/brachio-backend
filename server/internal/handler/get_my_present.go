package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetMyPresentsHandler struct {
	getMyPresentsUsecase usecase.GetMyPresentsInputPort
}

func (h *GetMyPresentsHandler) GetMyPresents(c echo.Context) error {
	userID := middleware.GetUserID(c)
	presents, err := h.getMyPresentsUsecase.Execute(c.Request().Context(), userID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := make([]*schema.Present, 0, len(presents))
	for _, p := range presents {
		resp = append(resp, &schema.Present{
			Id:        p.PresentID.String(),
			ItemCount: p.ItemCount,
			Message:   p.Message,
			Item: schema.ItemBase{
				Id:       p.Item.ItemID.String(),
				Name:     p.Item.Name,
				ImageUrl: p.Item.ImageURL,
			},
		})
	}

	return c.JSON(http.StatusOK, resp)
}
