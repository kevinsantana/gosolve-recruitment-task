include .env
export

SHELL=/bin/bash

all: help

envvars:
	export $(grep -v '^#' .env | xargs)

lint: ## lint: Apply golint.
	golangci-lint run -E gosec -E gofmt -E goimports --skip-dirs tests

help: ## help: Show this help message.
	@echo "usage: make [target] ..."
	@echo ""
	@echo "targets:"
	@grep -Eh '^.+:(\w+)?\ ##\ .+' ${MAKEFILE_LIST} | cut -d ' ' -f '3-' |  column -t -s ':' | egrep --color '^[^ ]*'

build: ## build: Build go executable
	go build .

run: ## run: run api
	go run . api