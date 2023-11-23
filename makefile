GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BINARY_NAME=ports
MAIN_PATH=cmd/ports/main.go

all: lint test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) ./...

lint:
	$(GOCMD) vet $(MAIN_PATH)
