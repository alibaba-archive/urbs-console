.PHONY: dev test mock doc

APP_NAME := urbs-console
APP_VERSION := $(shell git describe --tags --always --match "v[0-9]*")
APP_PATH := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

doc:
    # https://github.com/Mermade/widdershins
	# Install: npm i -g widdershins
	echo "# Content generated by 'make doc'. DO NOT EDIT.\n" > doc/swagger.yaml
	cat doc/swagger_header.yaml >> doc/swagger.yaml
	cat doc/swagger_user.yaml >> doc/swagger.yaml
	widdershins --language_tabs 'shell:Shell' 'http:HTTP' --summary doc/swagger.yaml -o doc/api.md

dev:
	@CONFIG_FILE_PATH=${PWD}/config/dev.yml APP_ENV=development go run main.go

test: 
	@CONFIG_FILE_PATH=${PWD}/config/test.yml STATIC_FILE_PATH=${PWD}/static APP_ENV=test go test ./...

mock:
	mockgen -source=./src/service/urbs_setting_interface.go -destination=./src/service/mock_service/urbs_setting__mock.go

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	@misspell -error $(GO_FILES)

.PHONY: coverhtml
coverhtml:
	@mkdir -p coverage
	@CONFIG_FILE_PATH=${PWD}/config/test.yml go test -coverprofile=coverage/cover.out ./...
	@go tool cover -html=coverage/cover.out -o coverage/coverage.html
	@go tool cover -func=coverage/cover.out | tail -n 1
