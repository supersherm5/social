MIGRATION_PATH = ./cmd/migrate/migrations
DB_URL = postgres://admin:adminpassword@localhost/social?sslmode=disable

#############
# APP
#############

# Build the app
run-app:
	@docker-compose up --build
.PHONY: run-app

# Stop the app
stop-app:
	@docker-compose down
.PHONY: stop-app


#############
# Migrations
#############

# Fix for Dirty migration version 3. Fix and force version.
fix-migration:
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) force 3
.PHONY: fix-migration

# Create a new migration
create-migration:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))
.PHONY: create-migration

# Run the migrations
migrate-up:
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) up
.PHONY: migrate-up

# Rollback the migrations
migrate-down:
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) down $(filter-out $@,$(MAKECMDGOALS))
.PHONY: migrate-down

#############
# Server
#############

# Build the server
build-server:
	@go build -o bin/main cmd/api/*.go
.PHONY: build-server

# Run the server
run-server:
	@go run cmd/api/*.go
.PHONY: run-server

# Run the server with hot reload
run-dev-server:
	@go run github.com/air-verse/air@v1.60.0 \
		--build.bin "./bin/main" \
		--build.cmd "make build-server" \
		--build.delay "100" \
		--build.exclude_dir "assets, bin, vender, testdata, docs, scripts" \
		--build.exclude_regex "*\\_test\\.go" \
		--build.include_ext "go, tpl, tmpl, html, css, xml, yaml, yml, json" \
		--build.log "build-errors.log" \
		--color.app "blue" \
		--color.build "yellow" \
		--color.main "magenta" \
		--color.runner "green" \
		--color.watcher "cyan" \
		--misc.clean_on_exit "true" \
.PHONY: run-dev-server


#############
# Utils
#############

# Format the code
fmt:
	@go fmt ./...
.PHONY: fmt