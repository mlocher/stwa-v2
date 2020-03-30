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
	# there is an issue with GOPATH if using the below command
	#$(MAKE) -C functions all
	cd functions; mkdir -p dist; go get ./...; go build -o dist/stwa ./...


build: clean functions
	yarn build
