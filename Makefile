.PHONY: tidy run test cover docker up down down-v logs

tidy:
	go mod tidy

run:
	go run ./cmd/atletismo-api

test:
	go test ./...

cover:
	go test ./internal/... -cover

docker:
	docker build -t atletismo-api .

up:
	docker compose up --build

down:
	docker compose down

down-v:
	docker compose down -v

logs:
	docker compose logs -f api