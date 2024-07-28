.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golangci-lint run ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build -o bin/todo cmd/todo/*.go
.PHONY: build

clean:
	rm -f bin/todo
.PHONY: clean

deps:
	go mod tidy
.PHONY: deps

