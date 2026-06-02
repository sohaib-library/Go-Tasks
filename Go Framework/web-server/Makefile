DB_URL=postgres://postgres:newpassword123@localhost:5432/postgres?sslmode=disable
MIGRATION_DIR=./migration/migrations

migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" status

create-migration:
	goose -dir $(MIGRATION_DIR) create $(name) sql
run:
	go run main.go