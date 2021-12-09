# ==============================================================================
# HELPERS
# ==============================================================================

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==============================================================================
# QUALITY CONTROL
# ==============================================================================

## audit: tidy depdendencies and format, vet, and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor depdendencies
.PHONY: vendor
vendor:
	@echo 'Tidying verifying module depdenencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies'
	go mod vendor	

# ==============================================================================
# BUILD
# ==============================================================================

current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/cli: build the cmd/cli application
.PHONY: build/cli
build/cli:
	@echo 'Building cmd/cli...'
	go build -ldflags=${linker_flags} -o=./bin/lsys ./cmd/cli
