# "./..." run command on everything in current directory and all directories inside.
.PHONY: build test clean fmt vet

build:
	go build ./...

test:
	go test -v -race ./...

clean:
	go clean ./...
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

all: fmt vet build test
