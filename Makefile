# Makefile
.SILENT:

# env var
include .env
export

# DSN
DSN := "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE)"

# Intall goose
goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

#==================================-MAIN-=====================================

G-push:
	git push -u origin tmozzze/feature

G-add:
	git add .

# Run docker-compose
run:
	docker-compose up -d

# Down docker-compose
down:
	docker-compose down

# Down and clean docker-compose
down-and-clean:
	docker-compose down -v

# Create new migration (make create-migration NAME=name)
create-migration:
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql
	@echo "Created migration: $(NAME)"

# Apply migrations
migrate-up:
	@echo "Applying migrations..."
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) up

# Rollback last migration
migrate-down:
	@echo "Rolling back last migration..."
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) down

# Show migration status
migrate-status:
	@echo "Migration status:"
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) status

# Rollback all migrations
migrate-reset:
	@echo "Rolling back migrations..."
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) reset