package main

import (
	"log"

	"github.com/yamato0211/brachio-backend/internal/infra/server"
)

func main() {
	server, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
