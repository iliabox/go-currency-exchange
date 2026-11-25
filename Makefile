.PHONY: build up down test migrate cron

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

purge:
	docker compose down -v

test:
	docker run --rm -v $(PWD):/src -w /src golang:1.27.4-alpine sh -c "go mod tidy && go test ./..."

migrate: build up
	docker run --rm --network go-exchange-rate_default -e DATABASE_URL="postgres://symfony:pswd@db:5432/app?sslmode=disable" go-exchange-rate ./migrate

cron: build up
	docker run --rm --network go-exchange-rate_default -e DATABASE_URL="postgres://symfony:pswd@db:5432/app?sslmode=disable" go-exchange-rate ./cron
