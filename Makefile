.PHONY: all test build

all: test
build:
		@gb build

install:
		@gb install

test:
		@gb test
