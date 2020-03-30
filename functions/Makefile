SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: all clean prepare stwa

all: stwa

clean:
	rm -rf dist/*

prepare: clean
	go get ./...

stwa: prepare
	go build -o dist/stwa ./...
