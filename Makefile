.PHONY: lint
lint:
	cd server && go tool golangci-lint run ./...
	cd ..

.PHONY: lint-fix
lint-fix:
	cd server && go tool golangci-lint run --fix ./...
	cd ..

.PHONY: docker_up
docker_up:
	docker-compose up -d --build

.PHONY: docker_down
docker_down:
	docker-compose down