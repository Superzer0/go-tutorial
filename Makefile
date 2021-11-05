all: test vet build 

test:
	go test ./...

vet: 
	go vet ./...

build:
	go build -o bin/worker ./cmd/worker 

run: 
	go run ./cmd/worker 
	