SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: clean serve functions build preview

clean:
	rm -rf dist/*

serve: 
	yarn serve

functions:
	cd functions; mkdir -p dist; go get ./...; go build -o dist/stwa ./...

build: clean functions
	yarn build
