# Makefile for Spotify Playlist Generator project

# Define Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=playlistgen
BINARY_UNIX=$(BINARY_NAME)_unix

# Define your .PHONY targets
.PHONY: all build test clean run deps

# Default command when you just run `make`
all: test build

# Build binary for current system
build:
	@echo "Building..."
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v ./cmd/...

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./pkg/...

# Clean up binaries
clean:
	@echo "Cleaning..."
	$(GOCMD) clean
	rm -f $(GOBIN)/$(BINARY_NAME)
	rm -f $(GOBIN)/$(BINARY_UNIX)

# Run the project
run:
	@echo "Running application..."
	$(GORUN) ./cmd/main.go

# Install dependencies
deps:
	@echo "Checking and downloading dependencies..."
	$(GOCMD) mod tidy
	$(GOCMD) mod download
