package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/brachio-backend/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/guregu/dynamo/v2"
)

type User struct {
	ID   int    `dynamo:"UserID,hash"`
	Name string `dynamo:"Name,range"`
	Age  int    `dynamo:"Age"`
	Text string `dynamo:"Text"`
}

var db *dynamo.DB

func init() {
	ctx := context.Background()
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db = dynamo.New(aws.Config{
		Region:       cfg.Dynamo.Region,
		BaseEndpoint: aws.String(cfg.Dynamo.Endpoint),
		Credentials:  credentials.NewStaticCredentialsProvider("dummy", "dummy", "dummy"),
	})

	// create users tables
	if err := db.CreateTable("Users", User{}).Run(ctx); err != nil {
		log.Fatal(err)
	}

	// put item
	u := User{
		ID:   1,
		Name: "John Doe",
		Age:  25,
		Text: "Hello, World!",
	}
	if err := db.Table("Users").Put(u).Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// get all items
		var results []User
		err := db.Table("Users").Scan().All(c.Request().Context(), &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, results)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
