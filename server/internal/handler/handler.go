package handler

import (
	"context"
	"log"
	"net/http"

	gorillawebsocket "github.com/gorilla/websocket"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/gateway/db"
	"github.com/yamato0211/brachio-backend/internal/gateway/memdb"
	websocket "github.com/yamato0211/brachio-backend/internal/gateway/pusher"
	"github.com/yamato0211/brachio-backend/internal/infra/dynamo"
	"github.com/yamato0211/brachio-backend/internal/usecase"
	pkgwebsocket "github.com/yamato0211/brachio-backend/pkg/websocket"
)

type Handler struct {
	// Card
	GetMyCardListHandler

	// Deck
	GetDeckListHandler
	PostDeckHandler
	GetDeckHandler
	PutDeckHandler
	DeleteDeckHandler

	// Item
	GetMyItemListHandler

	// Gacha
	GetGachaPowerHandler
	GetGachaListHandler
	PostGachaDrawHandler

	// Game
	GetWebSocketHandler

	// HealthCheck
	GetHealthCheckHandler

	// Present
	GetMyPresentsHandler
	ReceivePresentHandler

	// User
	GetUserHandler
}

func New() *Handler {
	ctx := context.Background()
	// DI
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	dc, err := dynamo.New(ctx, cfg.Dynamo)
	if err != nil {
		log.Fatal(err)
	}

	// repo
	masterCardRepo := db.NewMasterCardRepository(dc)
	deckRepo := db.NewDeckRepository(dc)
	masterItemRepo := db.NewMasterItemRepository(dc)
	userRepo := db.NewUserRepository(dc)
	presentRepo := db.NewPresentRepository(dc)
	gameStateRepo := memdb.NewGameStateRepository()

	// service
	pusher := pkgwebsocket.NewPusher()
	gameEventSender := websocket.NewGameEventSender(pusher)
	gameMasterService := service.NewGameMasterService(gameStateRepo, gameEventSender)
	abilityApplier := service.NewAbilityApplier()
	// skillApplier := service.NewSkillApprier()
	goodsApplier := service.NewGoodsApplier(gameMasterService)
	supporterApplier := service.NewSupporterApplier(gameMasterService)
	matcher := service.NewMatcherService()

	// usecase
	drawGachaUsecase := usecase.NewDrawGachaUsecase(masterCardRepo, userRepo, masterItemRepo)
	getMyDecksUsecase := usecase.NewGetMyDecksUsecase(deckRepo, masterCardRepo)
	getMyItemsUsecase := usecase.NewGetMyItemsUsecase(masterItemRepo, userRepo)
	createMyDeckUsecase := usecase.NewCreateMyDeckUsecase(deckRepo)
	getMyDeckUsecase := usecase.NewGetMyDeckUsecase(deckRepo, masterCardRepo)
	updateMyDeckUsecase := usecase.NewUpdateMyDeckUsecase(deckRepo)
	deleteMyDeckUsecase := usecase.NewDeleteMyDeckUsecase(deckRepo)
	getMyCardsUsecase := usecase.NewGetMyCardsUsecase(masterCardRepo, userRepo)
	getMyPresentUsecase := usecase.NewGetMyPresentsUsecase(presentRepo, masterItemRepo)
	receivePresentUsecase := usecase.NewReceivePresentUsecase(presentRepo, userRepo)
	getUserUsecase := usecase.NewGetUserUsecase(userRepo)

	// websocket usecases
	applyAbilityUsecase := usecase.NewApplyAbilityUsecase(gameStateRepo, abilityApplier)
	completeInitialPlacementUsecase := usecase.NewCompleteInitialPlacementUsecase(gameMasterService)
	evoluteMonsterUsecase := usecase.NewEvoluteMonsterUsecase(gameStateRepo, gameEventSender)
	flipCoinUsecase := usecase.NewFlipCoinUsecase(gameStateRepo, gameEventSender)
	giveUpUsecase := usecase.NewGiveUpUsecase(gameStateRepo)
	matchingUsecase := usecase.NewMatchingUsecase(gameStateRepo, deckRepo, matcher, gameMasterService)
	putInitializeMonsterUsecase := usecase.NewPutInitializeMonsterUsecase(gameStateRepo, gameMasterService)
	retreatUsecase := usecase.NewRetreatUsecase(gameStateRepo, gameEventSender)
	summonUsecase := usecase.NewSummonUsecase(gameStateRepo, gameEventSender)
	supplyEnergyUsecase := usecase.NewSupplyEnergyUsecase(gameStateRepo)
	useGoodsUsecase := usecase.NewUseGoodsUsecase(gameStateRepo, goodsApplier)
	useSupporterUsecase := usecase.NewUseSupporterUsecase(gameStateRepo, supporterApplier)

	return &Handler{
		GetMyCardListHandler: GetMyCardListHandler{
			getMyCardsUsecase: getMyCardsUsecase,
		},
		GetDeckListHandler: GetDeckListHandler{
			getMyDecksUsecase: getMyDecksUsecase,
		},
		PostDeckHandler: PostDeckHandler{
			createMyDeckUsecase: createMyDeckUsecase,
		},
		GetDeckHandler: GetDeckHandler{
			getMyDeckUsecase: getMyDeckUsecase,
		},
		PutDeckHandler: PutDeckHandler{
			updateMyDeckUsecase: updateMyDeckUsecase,
		},
		DeleteDeckHandler: DeleteDeckHandler{
			DeleteMyDeckUsecase: deleteMyDeckUsecase,
		},
		GetMyItemListHandler: GetMyItemListHandler{
			getMyItemsUsecase: getMyItemsUsecase,
		},
		GetGachaPowerHandler: GetGachaPowerHandler{},
		GetGachaListHandler:  GetGachaListHandler{},
		PostGachaDrawHandler: PostGachaDrawHandler{
			drawGachaUsecase: drawGachaUsecase,
		},
		GetHealthCheckHandler: GetHealthCheckHandler{},
		GetMyPresentsHandler: GetMyPresentsHandler{
			getMyPresentsUsecase: getMyPresentUsecase,
		},
		ReceivePresentHandler: ReceivePresentHandler{
			receivePresentUsecase: receivePresentUsecase,
		},
		GetUserHandler: GetUserHandler{
			getUserUsecase: getUserUsecase,
		},
		GetWebSocketHandler: GetWebSocketHandler{
			ApplyAbilityInputPort:             applyAbilityUsecase,
			CompleteInitialPlacementInputPort: completeInitialPlacementUsecase,
			EvoluteMonsterInputPort:           evoluteMonsterUsecase,
			FlipCoinInputPort:                 flipCoinUsecase,
			GiveUpInputPort:                   giveUpUsecase,
			MatchingInputPort:                 matchingUsecase,
			PutInitializeMonsterInputPort:     putInitializeMonsterUsecase,
			RetreatInputPort:                  retreatUsecase,
			SummonInputPort:                   summonUsecase,
			SupplyEnergyInputPort:             supplyEnergyUsecase,
			UseGoodsInputPort:                 useGoodsUsecase,
			UseSupporterInputPort:             useSupporterUsecase,
			upgrader: gorillawebsocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	}
}
