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

	dc, err := dynamo.New(context.Background(), cfg.Dynamo)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := db.NewUserRepository(dc)
	findUserUsecase := usecase.NewFindUserUsecase(userRepo)
	storeUserUsecase := usecase.NewStoreUserUsecase(userRepo)
	authMiddleware := middleware.NewAuthMiddleware(cfg, findUserUsecase, storeUserUsecase)

	e.Use(authMiddleware.Verify)
	e.Use(echomiddleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

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
