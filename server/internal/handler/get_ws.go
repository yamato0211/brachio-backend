package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	wsmsg "github.com/yamato0211/brachio-backend/internal/handler/schema/websocket"
	"github.com/yamato0211/brachio-backend/internal/usecase"
	"google.golang.org/protobuf/proto"
)

type GetWebSocketHandler struct {
	upgrader websocket.Upgrader

	ApplyAbilityInputPort         usecase.ApplyAbilityInputPort
	ApplyEnergyInputPort          usecase.ApplyEnergyInputPort
	AttackInputPort               usecase.AttackInputPort
	EvoluteMonsterInputPort       usecase.EvoluteMonsterInputPort
	GiveUpInputPort               usecase.GiveUpInputPort
	MatchingInputPort             usecase.MatchingInputPort
	PutInitializeMonsterInputPort usecase.PutInitializeMonsterInputPort
	RetreatInputPort              usecase.RetreatInputPort
	SummonInputPort               usecase.SummonInputPort
	UseGoodsInputPort             usecase.UseGoodsInputPort
	UseSupporterInputPort         usecase.UseSupporterInputPort
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

	ctx := c.Request().Context()
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

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to read message").SetInternal(err)
		}
		var req wsmsg.EventEnvelope
		if err := proto.Unmarshal(msg, &req); err != nil {
			return err
		}

		switch req := req.GetEvent().(type) {
		case *wsmsg.EventEnvelope_CreateRoomEventToServer:
			input := &usecase.MatchingInput{
				Password: req.CreateRoomEventToServer.Payload.Password,
			}
			h.MatchingInputPort.Execute(ctx, input)

		case *wsmsg.EventEnvelope_AbilityEventToServer:
		case *wsmsg.EventEnvelope_AttackMonsterEventToServer:
		case *wsmsg.EventEnvelope_CoinTossEventToServer:
		case *wsmsg.EventEnvelope_ConfirmEnergyEventToServer:
		case *wsmsg.EventEnvelope_ConfirmTargetEventToServer:
		case *wsmsg.EventEnvelope_DrawEventToServer:
		case *wsmsg.EventEnvelope_EvolutionMonsterEventToServer:
		case *wsmsg.EventEnvelope_InitialPlacementCompleteEventToServer:
		case *wsmsg.EventEnvelope_RetreatEventToServer:
		case *wsmsg.EventEnvelope_SummonMonsterEventToServer:
		case *wsmsg.EventEnvelope_SupplyEnergyEventToServer:
		case *wsmsg.EventEnvelope_SurrenderEventToServer:
		case *wsmsg.EventEnvelope_TakeGoodsEventToServer:
		case *wsmsg.EventEnvelope_TakeSupportEventToServer:
		case *wsmsg.EventEnvelope_EnterRoomEventToServer:
		case *wsmsg.EventEnvelope_ExchangeDeckEventToServer:
		}
	}
}
