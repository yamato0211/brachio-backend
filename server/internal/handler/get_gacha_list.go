package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/handler/schema"
)

type GetGachaListHandler struct{}

func (h *GetGachaListHandler) GetGachaList(c echo.Context) error {
	packs := []*schema.Gacha{
		{
			Id:       lo.ToPtr("1"),
			Name:     "最強の技術者達",
			ImageUrl: "https://pokepoke.kurichi.dev/images/kizuku-piece.avif",
		},
	}
	return c.JSON(http.StatusOK, packs)
}
