REQUIRED_BINS := go templ air npx
$(foreach bin,$(REQUIRED_BINS),\
    $(if $(shell command -v $(bin) 2> /dev/null),,$(error Please install `$(bin)`)))
	
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

update-templ:
	go install github.com/a-h/templ/cmd/templ@latest
.PHONY: update-templ
