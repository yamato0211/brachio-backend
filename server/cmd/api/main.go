package main

import (
	"log"

	"github.com/yamato0211/brachio-backend/internal/infra/server"
)

type User struct {
	ID   int    `dynamo:"UserID,hash"`
	Name string `dynamo:"Name,range"`
	Age  int    `dynamo:"Age"`
	Text string `dynamo:"Text"`
}

func main() {
	server, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
