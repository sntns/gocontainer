.PHONY: build test clean install lint help examples

# Default target
.DEFAULT_GOAL := help

# Variables
BINARY_NAME := gocontainer
BUILD_DIR := dist
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

## build: Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) ./cmd/gocontainer

## test: Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR) $(BINARY_NAME) coverage.out
	@rm -rf examples/*/container-output examples/*/{webserver,server,worker,configapp}

## install: Install the binary
install:
	@echo "Installing $(BINARY_NAME)..."
	@go install -ldflags="$(LDFLAGS)" ./cmd/gocontainer

## lint: Run linters
lint:
	@echo "Running linters..."
	@golangci-lint run

## examples: Build all examples
examples:
	@echo "Building examples..."
	@cd examples/webserver && go build -o webserver main.go
	@cd examples/multi-binary && go build -o server server.go && go build -o worker worker.go
	@cd examples/with-config && go build -o configapp main.go

## demo: Run a quick demo with the webserver example
demo: build examples
	@echo "Running demo..."
	@cd examples/webserver && ../../$(BINARY_NAME) build --binary webserver --label "app=webserver" --label "version=demo" --outdir ./container-output
	@echo "Demo complete! Check examples/webserver/container-output/"

## coverage: Generate test coverage report
coverage: test
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## deps: Download and verify dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod verify

## update: Update dependencies
update:
	@echo "Updating dependencies..."
	@go get -u all
	@go mod tidy

## help: Show this help
help:
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ { printf "  %-10s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
	@echo ""
	@echo "Variables:"
	@echo "  VERSION=$(VERSION)"
	@echo "  COMMIT=$(COMMIT)"
	@echo "  DATE=$(DATE)"