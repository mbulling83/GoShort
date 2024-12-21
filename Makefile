# Project variables
APP_NAME = goshort
BIN_DIR = bin
BUILD_DIR = $(BIN_DIR)/$(APP_NAME)
PKG_LIST = $(shell go list ./... | grep -v /tests)

# Commands
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOLINT = golangci-lint
GOFMT = gofmt
GOGET = $(GOCMD) get
GOMODTIDY = $(GOCMD) mod tidy
DOCKER_COMPOSE = docker-compose

# Build variables
BUILD_FLAGS = -ldflags "-s -w"
ENV_FILE = .env

# Targets
.PHONY: all build test lint fmt run clean install docker-up docker-down docker-build

all: build

build: clean
	@echo "Building the $(APP_NAME)..."
	$(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR) ./cmd/server
	@echo "Build completed."

test:
	@echo "Running tests..."
	$(GOTEST) -v -race ./...
	@echo "Tests completed."

lint:
	@echo "Running lint checks..."
	$(GOLINT) run ./...
	@echo "Linting completed."

fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w $(PKG_LIST)
	@echo "Code formatted."

run: build
	@echo "Running the application..."
	$(BUILD_DIR)

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	@echo "Cleanup completed."

install-deps:
	@echo "Installing dependencies..."
	$(GOGET) -u ./...
	$(GOMODTIDY)
	@echo "Dependencies installed."

migrate:
	@echo "Running database migrations..."
	go run ./cmd/migrate
	@echo "Migrations completed."

seed:
	@echo "Seeding the database..."
	bash scripts/seed.sh

cleanup:
	@echo "Cleaning up expired URLs..."
	bash scripts/cleanup.sh

docker-up:
	@echo "Starting Docker containers..."
	$(DOCKER_COMPOSE) up -d
	@echo "Docker containers started."

docker-down:
	@echo "Stopping Docker containers..."
	$(DOCKER_COMPOSE) down
	@echo "Docker containers stopped."

docker-build:
	@echo "Building Docker images..."
	$(DOCKER_COMPOSE) build
	@echo "Docker images built."
help:
	@echo "Available targets:"
	@echo "  build        Build the application"
	@echo "  test         Run unit tests"
	@echo "  lint         Run lint checks"
	@echo "  fmt          Format the code"
	@echo "  run          Build and run the application"
	@echo "  clean        Clean up the build artifacts"
	@echo "  install-deps Install project dependencies"
	@echo "  migrate      Run database migrations"
	@echo "  seed         Seed the database with test data"
	@echo "  cleanup      Remove expired URLs from the database"
	@echo "  docker-up    Start Docker containers (app and database)"
	@echo "  docker-down  Stop Docker containers"
	@echo "  docker-restart Restart Docker containers"
