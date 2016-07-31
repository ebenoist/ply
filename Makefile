.PHONY: all test build

name = ply

all: test build
build:
		@go build -o bin/ply ./cmd/ply

test:
		@go test

release:
	@env GOOS=linux go build
	@env GOOS=darwin go build
