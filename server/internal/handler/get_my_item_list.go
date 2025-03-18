package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type GetMyItemListHandler struct {
	getMyItemsUsecase usecase.GetMyItemsInputPort
}

func (h *GetMyItemListHandler) GetMyItemList(c echo.Context) error {
	userID := middleware.GetUserID(c)
	items, err := h.getMyItemsUsecase.Execute(c.Request().Context(), userID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := make([]*schema.Item, 0, len(items))
	for _, item := range items {
		resp = append(resp, &schema.Item{
			Id:    item.MasterItem.ItemID.String(),
			Name:  item.MasterItem.Name,
			Count: item.Count,
		})
	}
	return c.JSON(http.StatusOK, resp)
}
