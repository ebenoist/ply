.PHONY: all test build

name = ply

all: test
build:
		@gb build

test:
		@gb test

release:
	@env GOOS=linux gb build
	@env GOOS=darwin gb build
