$(shell cp -n .env.example .env)
include .env
export

all: air
.PHONY: all

test: 
	@go test -v ./...
.PHONY: test

air:
	@air
.PHONY: air
