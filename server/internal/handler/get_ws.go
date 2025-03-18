package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type GetWebSocketHandler struct {
	upgrader websocket.Upgrader
}

func (h *GetWebSocketHandler) Ws(c echo.Context) error {
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to upgrade connection").SetInternal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Warn("failed to close websocket", "error", err)
		}
	}()

	// Ping Pong
	go func() {
		ticker := time.NewTicker(50 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				break
			}
		}
	}()
	conn.SetPongHandler(func(string) error {
		return conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	})

	// TODO: マッチング開始

	// TODO:

	return nil
}
