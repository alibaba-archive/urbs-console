.PHONY: dev test

APP_NAME := urbs-console
APP_VERSION := $(shell git describe --tags --always --match "v[0-9]*")
APP_PATH := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

dev:
	@CONFIG_FILE_PATH=${PWD}/config/default.yml APP_ENV=development go run main.go

test:
	@CONFIG_FILE_PATH=${PWD}/config/test.yml APP_ENV=test go test -v ./...