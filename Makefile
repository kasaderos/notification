.PHONY: build run test lint clean docker-build docker-run migrate-up migrate-down

# Build the application
build:
	go build -o bin/tgbot ./cmd/tgbot
	go build -o bin/job ./cmd/job

# Run the telegram bot
run-bot: build
	./bin/tgbot

# Run the job service
run-job: build
	./bin/job

# Run both services
run: build
	@echo "Starting both services..."
	@echo "Run './bin/tgbot' in one terminal and './bin/job' in another"

# Run tests
test:
	go test ./...

# Run linter
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t notification-bot .

# Run Docker container
docker-run:
	docker run --env-file .env notification-bot

# Run database migrations up
migrate-up:
	# TODO: Implement migration runner
	@echo "Migrations would be run here"

# Run database migrations down
migrate-down:
	# TODO: Implement migration rollback
	@echo "Migrations would be rolled back here"

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Generate mocks (if using mockgen)
mocks:
	# TODO: Add mock generation if needed
	@echo "Mocks would be generated here"

# Setup development environment
setup: deps
	@echo "Setting up development environment..."
	@echo "Make sure to set up your .env file with required environment variables:"
	@echo "- TELEGRAM_BOT_TOKEN"
	@echo "- POSTGRES_DSN"
	@echo "- ELASTICSEARCH_URL"

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Build and run the application"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linter"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  migrate-up   - Run database migrations"
	@echo "  migrate-down - Rollback database migrations"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  setup        - Setup development environment"
	@echo "  help         - Show this help message"