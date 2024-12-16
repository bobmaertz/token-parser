# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

VERSION_PKG=github.com/bobmaertz/token-parser/pkg
COMMIT      := $(shell git rev-parse --short=10 HEAD)
BUILD_DATE  := $(shell date -u +%Y-%m-%d)
LDFLAGS     := -X $(VERSION_PKG)/version.Commit=$(COMMIT) -X $(VERSION_PKG)/version.BuildDate=$(BUILD_DATE)

# Main package path
MAIN_PATH=./cmd/token-parser

# Binary name
BINARY_NAME=./bin/token-parser

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "$(LDFLAGS)" -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "$(LDFLAGS)" -v $(MAIN_PATH)
	$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) -v $(MAIN_PATH)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BINARY_NAME).exe -v $(MAIN_PATH)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) --ldflags "$(LDFLAGS)" -o $(BINARY_NAME) -v $(MAIN_PATH)

.PHONY: all build test clean run build-linux build-windows build-mac