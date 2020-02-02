.DEFAULT_TARGET=help
.PHONY: all
all: help

# VARIABLES
APP_NAME = secure
GO_VERSION ?= 1.13
GO_FILES = $(shell go list ./... | grep -v /vendor/)
PROJECT_PATH ?= github.com/pypl-johan/secure

VERSION ?= $(shell git describe --exact-match --tags 2>/dev/null)
COMMIT = $(shell git rev-parse HEAD | cut -c 1-6)
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

LDFLAGS = -ldflags "-s -w -X ${PROJECT_PATH}/cmd/version.Version=v${VERSION} -X ${PROJECT_PATH}/cmd/version.Commit=${COMMIT} -X ${PROJECT_PATH}/cmd/version.BuildTime=${BUILD_TIME}"
RUN_CMD = docker run --rm -it -v "$(GOPATH):/go" -v "$(CURDIR)":/go/src/${PROJECT_PATH} -w /go/src/${PROJECT_PATH} golang:${GO_VERSION}-stretch

# COMMANDS

## lint: run linter
.PHONY: lint
lint:
	$(call blue, "# running linter...")
	@golangci-lint run -c .golangci.yml

## ci-test: run test suite for application
.PHONY: ci-test
ci-test:
	$(call blue, "# running tests...")
	@gotestsum --junitfile $(TEST_RESULTS_DIR)/unit-tests.xml -- \
		-race ./... \
		-coverprofile cover.out \
		-covermode atomic \
		-coverpkg ./...

## test: run test suite for application
.PHONY: test
test:
	$(call blue, "# running tests...")
	@${RUN_CMD} go test -cover -race ./...

## install: install tool to your $GOPATH/bin directory
.PHONY: install
install:
	$(call blue, "# installing...")
	@go build $(LDFLAGS) -o $(GOPATH)/bin/secure

## help: Show this help message
.PHONY: help
help: Makefile
	@echo "${APP_NAME} - v${VERSION}"
	@echo " Choose a command run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^## //p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

# FUNCTIONS
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef

