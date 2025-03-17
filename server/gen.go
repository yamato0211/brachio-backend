package server

//go:generate go tool oapi-codegen --config=oapi-schema.cfg.yaml ../brachio-api-spec/openapi/build/openapi.yaml
//go:generate go tool oapi-codegen --config=oapi-server.cfg.yaml ../brachio-api-spec/openapi/build/openapi.yaml
