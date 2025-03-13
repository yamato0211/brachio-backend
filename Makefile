.PHONY: lint
lint:
	go tool golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	go tool golangci-lint run --fix ./...

.PHONY: docker_up
docker_up:
	docker-compose up -d

.PHONY: docker_down
docker_down:
	docker-compose down