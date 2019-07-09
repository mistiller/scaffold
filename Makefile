# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=./app/app

all: deps test build-linux
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/goexp/main.go
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v ./cmd/goexp/main.go
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
deps:
	$(GOGET) github.com/go-redis/redis
	$(GOGET) github.com/boltdb/bolt
build-docker-compose: 
	docker-compose build && docker-compose up && make clean
docker:
	docker build -t goexp . && docker run goexp