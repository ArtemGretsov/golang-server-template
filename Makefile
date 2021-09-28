-include .env
-include .env.local

POSTGRES_CONNECT_STRING=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR=./src/database/migrations

app-test:
	@TEST_EVN=1 go test ./src/... | grep -v "no test files"

app-benchmark:
	@go test -bench Benchmark -benchtime=1000x -benchmem ./src/...

db-run:
	@docker-compose --env-file .env.local up -d

db-migrate-up:
	@go run src/database/migrations/migration.go up

db-migrate-down:
	@go run src/database/migrations/migration.go down

db-migrate-generate:
	@go run src/database/migrations/migration.go generate $(name)

ent-create:
	@go run entgo.io/ent/cmd/ent init --target ./src/database/schema $(name)

ent-generate:
	@rm -rf ./src/database/_schemagen
	@go run entgo.io/ent/cmd/ent generate --target ./src/database/_schemagen ./src/database/schema
