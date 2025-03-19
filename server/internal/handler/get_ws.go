package handler

import (
	"log/slog"
	"net/http"
	"time"

	// "github.com/gorilla/websocket"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	wsmsg "github.com/yamato0211/brachio-backend/internal/handler/schema/websocket"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/websocket/payload"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
	pkgws "github.com/yamato0211/brachio-backend/pkg/websocket"
	"google.golang.org/protobuf/proto"
)

type GetWebSocketHandler struct {
	upgrader websocket.Upgrader
	pusher   *pkgws.Pusher

	ApplyAbilityInputPort             usecase.ApplyAbilityInputPort
	SupplyEnergyInputPort             usecase.SupplyEnergyInputPort
	AttackInputPort                   usecase.AttackInputPort
	EvoluteMonsterInputPort           usecase.EvoluteMonsterInputPort
	GiveUpInputPort                   usecase.GiveUpInputPort
	MatchingInputPort                 usecase.MatchingInputPort
	PutInitializeMonsterInputPort     usecase.PutInitializeMonsterInputPort
	RetreatInputPort                  usecase.RetreatInputPort
	SummonInputPort                   usecase.SummonInputPort
	UseGoodsInputPort                 usecase.UseGoodsInputPort
	UseSupporterInputPort             usecase.UseSupporterInputPort
	CompleteInitialPlacementInputPort usecase.CompleteInitialPlacementInputPort
	FlipCoinInputPort                 usecase.FlipCoinInputPort
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

	rawUserID := middleware.GetUserID(c)
	userID, err := model.ParseUserID(rawUserID)
	if err != nil {
		return nil
	}

	runner, closer := h.pusher.Register(ctx, userID, conn)
	defer closer()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				slog.ErrorContext(ctx, "panic", "error", err)
			}
		}()

		if err := runner(ctx); err != nil {
			slog.ErrorContext(ctx, "JoinRoomHandler:runner", "error", err)
		}
	}()

	var roomID string
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
		case *wsmsg.EventEnvelope_EnterRoomEventToServer: // ルームに入るイベント
			input := &usecase.MatchingInput{
				UserID:   rawUserID,
				Password: req.EnterRoomEventToServer.Payload.Password,
				DeckID:   req.EnterRoomEventToServer.Payload.DeckId,
			}
			_roomID, err := h.MatchingInputPort.Execute(ctx, input)
			if err != nil {
				c.Logger().Warnf("failed to apply ability: %v", err)
			}
			roomID = _roomID

		case *wsmsg.EventEnvelope_InitialSummonEventToServer: // モンスターを初期配置するイベント
			input := &usecase.PutInitializeMonsterInput{
				RoomID:   roomID,
				UserID:   rawUserID,
				CardID:   req.InitialSummonEventToServer.Payload.Card.Id,
				Position: int(req.InitialSummonEventToServer.Payload.Position),
			}

			if err := h.PutInitializeMonsterInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to put initialize monster: %v", err)
			}

		case *wsmsg.EventEnvelope_InitialPlacementCompleteEventToServer: // 初期配置完了イベント
			input := &usecase.CompleteInitialPlacementInput{
				RoomID: roomID,
				UserID: rawUserID,
			}
			if err := h.CompleteInitialPlacementInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to complete initial placement: %v", err)
			}

		case *wsmsg.EventEnvelope_AbilityEventToServer: // アビリティを使うイベント
			input := &usecase.ApplyAbilityInput{
				RoomID: roomID,
				UserID: rawUserID,
			}
			if err := h.ApplyAbilityInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to apply ability: %v", err)
			}

		case *wsmsg.EventEnvelope_AttackMonsterEventToServer: // モンスターを攻撃するイベント
			input := &usecase.AttackInput{
				RoomID: roomID,
				UserID: rawUserID,
			}
			if err := h.AttackInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to attack: %v", err)
			}

		case *wsmsg.EventEnvelope_CoinTossEventToServer: // コイントスをしたよイベント
			input := &usecase.SupplyEnergyInput{
				RoomID: roomID,
				UserID: rawUserID,
			}
			if err := h.SupplyEnergyInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to apply energy: %v", err)
			}

		case *wsmsg.EventEnvelope_SummonMonsterEventToServer: // モンスターを召喚するイベント
			input := &usecase.SummonInput{
				RoomID:   roomID,
				UserID:   rawUserID,
				CardID:   req.SummonMonsterEventToServer.Payload.Card.Id,
				Position: int(req.SummonMonsterEventToServer.Payload.Position),
			}
			if err := h.SummonInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to summon: %v", err)
			}

		case *wsmsg.EventEnvelope_EvolutionMonsterEventToServer: // モンスターを進化させるイベント
			input := &usecase.EvoluteMonsterInput{
				RoomID:   roomID,
				UserID:   rawUserID,
				CardID:   req.EvolutionMonsterEventToServer.Payload.Card.Id,
				Position: int(req.EvolutionMonsterEventToServer.Payload.Position),
			}
			if err := h.EvoluteMonsterInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to evolute monster: %v", err)
			}

		case *wsmsg.EventEnvelope_RetreatEventToServer: // にげるイベント
			input := &usecase.RetreatInput{
				RoomID:    roomID,
				UserID:    rawUserID,
				RetreatTo: int(req.RetreatEventToServer.Payload.GetPosition()),
			}
			if err := h.RetreatInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to retreat: %v", err)
			}

		case *wsmsg.EventEnvelope_SupplyEnergyEventToServer: // エネルギーをつけるイベント
			input := &usecase.SupplyEnergyInput{
				RoomID: roomID,
				UserID: rawUserID,
				Positions: lo.Map(req.SupplyEnergyEventToServer.Payload.Supplys, func(s *payload.SupplyEnergys, _ int) []model.MonsterType {
					return lo.Map(s.Energies, func(e messages.Element, _ int) model.MonsterType {
						switch e {
						case messages.Element_ALCHOHOL:
							return model.MonsterTypeAlchohol
						case messages.Element_KNOWLEDGE:
							return model.MonsterTypeKnowledge
						case messages.Element_MONEY:
							return model.MonsterTypeMoney
						case messages.Element_MUSCLE:
							return model.MonsterTypeMuscle
						case messages.Element_POPULARITY:
							return model.MonsterTypePopularity
						case messages.Element_NULL:
							return model.MonsterTypeNull
						default:
							return model.MonsterTypeNull
						}
					})
				}),
			}
			if err := h.SupplyEnergyInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to apply energy: %v", err)
			}

		case *wsmsg.EventEnvelope_SurrenderEventToServer:
			input := &usecase.GiveUpInput{
				RoomID: roomID,
				UserID: rawUserID,
			}
			if err := h.GiveUpInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to give up: %v", err)
			}

		case *wsmsg.EventEnvelope_TakeGoodsEventToServer:
			input := &usecase.UseGoodsInput{
				RoomID: roomID,
				UserID: rawUserID,
				CardID: req.TakeGoodsEventToServer.Payload.Card.Id,
				// Targets: req.TakeGoodsEventToServer.Payload.Targets,
			}
			if err := h.UseGoodsInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to use goods: %v", err)
			}

		case *wsmsg.EventEnvelope_TakeSupportEventToServer:
			input := &usecase.UseSupporterInput{
				RoomID: roomID,
				UserID: rawUserID,
				CardID: req.TakeSupportEventToServer.Payload.Card.Id,
				// Targets:
			}
			if err := h.UseSupporterInputPort.Execute(ctx, input); err != nil {
				c.Logger().Warnf("failed to use supporter: %v", err)
			}
		}
	}
}
