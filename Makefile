FOLDER_MIGRATIONS = src/migrations
FOLDER_SERVER = src/server

install-tools:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

migrate-create:
	@go run $(FOLDER_SERVER)/main.go migrate-create ${name}

migrate:
	@go run $(FOLDER_SERVER)/main.go migrate

migrate-down:
	@go run $(FOLDER_SERVER)/main.go migrate-down

run:
	@go run $(FOLDER_SERVER)/main.go server

