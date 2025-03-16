package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `dynamo:"UserID,hash"`
	Name string `dynamo:"Name,range"`
	Age  int    `dynamo:"Age"`
	Text string `dynamo:"Text"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
