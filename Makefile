.PHONY: all test build

name = ply

all: test
build:
		@gb build

test:
		@gb test

release:
	@rm -rf dist/*
	@env GOOS=linux gb build
	@mv bin/ply dist/$(name)-linux.v${VERSION}
	@env GOOS=darwin gb build
	@mv bin/ply dist/$(name)-osx.v${VERSION}
