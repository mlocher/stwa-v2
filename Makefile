SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: all clean serve functions build preview

all: build

clean:
	rm -rf dist/*

serve:
	yarn serve

functions:
	$(MAKE) -C functions all

build: clean functions
	yarn build
