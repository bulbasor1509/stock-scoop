DB_DRIVER=postgres
DB_URL=postgres://root:root@localhost:5432/stock_data?sslmode=disable
MIGRATIONS_DIR=internal/sql/schema

.PHONY: migrate-up migrate-down migrate-status

migrate_up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up

migrate_down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down

migrate_status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" status