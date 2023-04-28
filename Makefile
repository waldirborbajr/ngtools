ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain=2>/dev/null))
  GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

VERSION_PACKAGE=localhost/ngtools/cmd

SYS_VERSION:="v0.1.0"

GO_LDFLAGS += \
	-w \
	"-extldflags=-static" \
  -X $(VERSION_PACKAGE).Version=$(SYS_VERSION) \
  -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
  -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
  -X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
  -X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: all
all: format build

.PHONY: build
build:
	go build -v -a -ldflags "$(GO_LDFLAGS)" -o bin/ngtools main.go

.PHONY: run
run: build
	./bin/ngtools start -protocol http -port 3031

.PHONY: tidy
tidy: 
	@go mod tidy

.PHONY: format
format: 
	@gofmt -s -w ./
