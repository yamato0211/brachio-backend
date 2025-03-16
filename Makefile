.PHONY: lint
lint:
	cd server && go tool golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	cd server && go tool golangci-lint run --fix ./...

.PHONY: docker_up
docker_up:
	docker-compose up -d --build

.PHONY: docker_down
docker_down:
	docker-compose down

gen/openapi:
	cd server && go tool oapi-codegen --o internal/handler/openapi/openapi.gen.go ../openapi/openapi.yaml