# Development
run-local:
	@echo "Running locally..."
	go run src/main.go

run:
	@echo "Running..."
	chmod +x bin/qotd
	bin/qotd

test:
	@echo "Running tests..."
	go test -v ./src/test/... ./src/internal/...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./test/... ./src/internal/...
	go test -coverprofile=coverage.out ./test/... ./src/internal/...
	go tool cover -html=coverage.out -o coverage.html

build:
	@echo "Building..."
	mkdir -p bin
	go build -o bin/qotd src/main.go

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Utility commands
clear-cache:
	@echo "Clearing quote cache..."
	curl -X POST http://localhost:8080/clear-cache

.PHONY: run-local test test-coverage build clean clear-cache
