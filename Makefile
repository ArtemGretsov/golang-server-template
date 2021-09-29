-include .env
-include .env.local

POSTGRES_CONNECT_STRING=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR=./internal/database/migrations

test:
	@TEST_EVN=1 go test ./internal/... | grep -v "no test files"

benchmark:
	@go test -bench Benchmark -benchtime=1000x -benchmem ./internal/...

env-run:
	@docker-compose --env-file .env.local up -d

migrate-up:
	@go run internal/database/migrations/migration.go up

migrate-down:
	@go run internal/database/migrations/migration.go down

migrate-generate:
	@go run internal/database/migrations/migration.go generate $(name)

ent-create:
	@go run entgo.io/ent/cmd/ent init --target ./internal/database/schema $(name)

ent-generate:
	@rm -rf ./internal/database/_schemagen
	@go run entgo.io/ent/cmd/ent generate --target ./internal/database/_schemagen ./internal/database/schema
