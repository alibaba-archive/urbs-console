.PHONY: dev test mock

APP_NAME := urbs-console
APP_VERSION := $(shell git describe --tags --always --match "v[0-9]*")
APP_PATH := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

dev:
	@CONFIG_FILE_PATH=${PWD}/config/default.yml APP_ENV=development go run main.go

test:
	@CONFIG_FILE_PATH=${PWD}/config/default.yml APP_ENV=test go test -v ./...

mock:
	mockgen -source=./src/service/urbs_setting_interface.go -destination=./src/service/mock_service/urbs_setting__mock.go