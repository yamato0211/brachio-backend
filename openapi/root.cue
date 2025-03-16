package root

import (
	"github.com/yamato0211/brachio-backend/openapi/path"
	"github.com/yamato0211/brachio-backend/openapi/definition"
	"github.com/yamato0211/brachio-backend/openapi/component/schema"
)

definition.#openapi & {
	info: {
		title:   "The PokePoke API"
		version: "0.0.1"
	}
	servers: [
		{
			url:         "http://localhost:3000"
			description: "local server"
		},
	]
	paths: path
	components: {
		schemas: schema
		securitySchemes: {
      bearerAuth: {
        type: "http"
        scheme: "bearer"
        bearerFormat: "JWT"
      }
		}
	}
}