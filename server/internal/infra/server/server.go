package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/gateway/db"
	"github.com/yamato0211/brachio-backend/internal/handler"
	"github.com/yamato0211/brachio-backend/internal/infra/cognito"
	"github.com/yamato0211/brachio-backend/internal/infra/dynamo"
	"github.com/yamato0211/brachio-backend/internal/infra/middleware"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

type Server struct {
	e    *echo.Echo
	port int
}

func New() (*Server, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil

	e := echo.New()

	ctx := context.Background()

	dc, err := dynamo.New(ctx, cfg.Dynamo)
	if err != nil {
		log.Fatal(err)
	}

	cc, err := cognito.New(ctx, cfg.Cognito)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := db.NewUserRepository(dc)
	findUserUsecase := usecase.NewFindUserUsecase(userRepo)
	storeUserUsecase := usecase.NewStoreUserUsecase(userRepo)
	authMiddleware := middleware.NewAuthMiddleware(cfg, findUserUsecase, storeUserUsecase, cc)

	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Accept", "Origin"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	e.Use(authMiddleware.Verify)

	e.GET("/", HealthCheck)

	handler := handler.New()
	RegisterHandlers(e, handler)

	return &Server{
		e:    e,
		port: cfg.Server.Port,
	}, nil
}

func (s *Server) Run() error {
	return s.e.Start(fmt.Sprintf(":%d", s.port))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
