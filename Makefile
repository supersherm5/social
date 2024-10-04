


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