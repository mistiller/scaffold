# Go parameters
all: run
test:
	go test -v ./...
build: 
	docker build -t goexp .
run:
	docker build -t goexp . && docker run goexp
run-compose:
	docker-compose build && docker-compose up