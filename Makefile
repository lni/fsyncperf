MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
.DELETE_ON_ERROR:
.SUFFIXES:

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build  -o bin/fsyncpref   main.go
