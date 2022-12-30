VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')
DOCKER := $(shell which docker)

export GO111MODULE = on

all: ci-lint ci-test install

###############################################################################
# Build / Install
###############################################################################

LD_FLAGS = -X github.com/forbole/juno/v3/cmd.Version=$(VERSION) \
	-X github.com/forbole/juno/v3/cmd.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building autox binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/autox.exe ./cmd/autox
else
	@echo "building autox binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/autox ./cmd/autox
endif

install: go.sum
	@echo "installing autox binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/autox

proto-gen:
	@scripts/protocgen.sh

###############################################################################
# Tests / CI
###############################################################################

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" | xargs goimports -w -local github.com/saiSunkari19/psql-juno
.PHONY: format

clean:
	rm -f tools-stamp ./build/**

.PHONY: install build ci-test ci-lint clean