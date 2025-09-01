.PHONY: all test build clean lint security

# Variables
BINARY_NAME=helmgraph

all: test build clean lint security

## test - Run all go tests
test:
	@echo "Running tests..."
	go test -v ./...

## build - Build the go binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) ./cmd/helmgraph

lint: go-lint

## clean - Remove the built binary
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

## go-lint - Run go vet and go fmt
go-lint:
	@echo "Running go vet..."
	go vet ./...
	@echo "Running go fmt..."
	go fmt ./...

## security - Run trivy filesystem scan
security:
	@echo "Running trivy filesystem scan..."
	trivy fs .
