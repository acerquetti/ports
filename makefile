GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BINARY_NAME=ports
MAIN_PATH=cmd/ports/main.go
DOCKER_BUILD_GO_ENV=CGO_ENABLED=0 GOARCH=amd64 GOHOSTARCH=amd64 GOHOSTOS=linux GOOS=linux

all: lint test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) ./...

lint:
	$(GOCMD) vet $(MAIN_PATH)

docker: lint test
	$(DOCKER_BUILD_GO_ENV) $(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)
	docker build -t $(BINARY_NAME) .
