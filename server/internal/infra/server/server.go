package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/handler"
	"github.com/yamato0211/brachio-backend/internal/infra/dynamo"
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

	dc, err := dynamo.New(context.Background(), cfg.Dynamo)
	if err != nil {
		slog.Warn("failed to create dynamo client", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
	}

	e.GET("/", HealthCheck)

	e.GET("/users", func(c echo.Context) error {
		type User struct {
			ID   string `dynamo:"UserId,hash"`
			Name string `dynamo:"Name"`
		}
		users := []User{}
		err := dc.Table("Users").Scan().All(c.Request().Context(), &users)
		if err != nil {
			slog.Error("error!", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, users)
	})

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
