package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/handler"
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
