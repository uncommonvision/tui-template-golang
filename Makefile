# TUI Template Makefile

BINARY_NAME=mytui
VERSION=1.0.0
BUILD_DIR=build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X tui-template/cmd.version=$(VERSION)"

.PHONY: all build build-all clean test deps dev run help

all: help

# Build the binary
build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) .

# Build for multiple platforms
build-all: clean
	mkdir -p $(BUILD_DIR)
	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .
	# macOS  
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run the application
run: build
	./$(BINARY_NAME)

# Run main TUI (same as run since it's now the default)
start: build
	./$(BINARY_NAME)

# Run examples
examples-list: build
	./$(BINARY_NAME) examples list

examples-form: build
	./$(BINARY_NAME) examples form

# Initialize config
config-init: build
	./$(BINARY_NAME) config init

# Generate documentation
docs:
	$(GOCMD) doc -all ./...

# Install the binary to $GOROOT/bin
install: build
	cp $(BINARY_NAME) $(GOROOT)/bin/

# Uninstall the binary
uninstall:
	rm -f $(GOROOT)/bin/$(BINARY_NAME)

# Show help
help:
	@echo "Available commands:"
	@echo "  build         Build the binary"
	@echo "  build-all     Build for multiple platforms"
	@echo "  clean         Clean build artifacts"
	@echo "  test          Run tests"
	@echo "  deps          Install dependencies"
	@echo "  run           Build and run the application"
	@echo "  start         Build and run the main TUI"
	@echo "  examples-*    Build and run example applications"
	@echo "  config-init   Initialize configuration file"
	@echo "  docs          Generate documentation"
	@echo "  install       Install binary to GOROOT/bin"
	@echo "  uninstall     Remove binary from GOROOT/bin"
	@echo "  help          Show this help message"
