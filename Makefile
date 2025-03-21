build: gen-swag
	@go build -o bin/online_library cmd/app/main.go

gen-swag:
	@swag init -g ./cmd/app/main.go

run:build
	@./bin/online_library

new-migrate:
	@migrate create -ext sql -dir db/migrations -seq ${name}

ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

DATABASE_URL := "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate-up:
	@migrate -database "$(DATABASE_URL)" -path db/migrations up

migrate-down:
	@migrate -database "$(DATABASE_URL)" -path db/migrations down 1