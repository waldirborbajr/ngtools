.PHONY: build clean contrib_check coverage docker-build docker-install help install isntall lint run size test uninstall

# detect GOPATH if not set
ifndef $(GOPATH)
	$(info GOPATH is not set, autodetecting..)
	TESTPATH := $(dir $(abspath ../../..))
	DIRS := bin pkg src

	# create a ; separated line of tests and pass it to shell
	MISSING_DIRS := $(shell $(foreach entry,$(DIRS),test -d "$(TESTPATH)$(entry)" || echo "$(entry)";))
	ifeq ($(MISSING_DIRS),)
		$(info Found GOPATH: $(TESTPATH))
		export GOPATH := $(TESTPATH)
	else
		$(info ..missing dirs "$(MISSING_DIRS)" in "$(TESTDIR)")
		$(info GOPATH autodetection failed)
	endif
endif

# Set go modules to on and use GoCenter for immutable modules
# export GO111MODULE = on
# export GOPROXY = https://proxy.golang.org,direct

# Determines the path to this Makefile
THIS_FILE := $(lastword $(MAKEFILE_LIST))

GOBIN := $(GOPATH)/bin

APP=ngtools

define HEADER
                __                .__          
  ____    _____/  |_  ____   ____ |  |   ______
 /    \  / ___\   __\/  _ \ /  _ \|  |  /  ___/
|   |  \/ /_/  >  | (  <_> |  <_> )  |__\___ \ 
|___|  /\___  /|__|  \____/ \____/|____/____  >
     \//_____/                              \/

endef
export HEADER

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
	-s \
	-w \
	'-extldflags=-static' \
  -X $(VERSION_PACKAGE).Version=$(SYS_VERSION) \
  -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
  -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
  -X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
  -X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

## build: builds a local version
build:
	@echo "$$HEADER"
	@echo "Building..."
	go build -a -ldflags "$(GO_LDFLAGS)" -o bin/${APP} main.go
	@echo "Done building"

## tidy: runs tidy
tidy: 
	@go mod tidy

## format: runs gofmt
format: 
	@gofmt -s -w ./

## clean: removes old build cruft
clean:
	rm -rf ./dist
	rm -rf ./bin/${APP}
	@echo "Done cleaning"

## test: runs the test suite
test: build
	@echo "$$HEADER"
	go test -v -timeout=1s -race -covermode=atomic -count=1 ./...

## coverage: figures out and displays test code coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

## gosec: runs the gosec static security scanner against the source code
gosec: $(GOBIN)/gosec
	gosec -tests ./...

$(GOBIN)/gosec:
	cd && go install github.com/securego/gosec/v2/cmd/gosec@latest

## run: execute binary version
run: build
	./bin/ngtools start -protocol http -port 3031

## update: update go.mod
update:
	go get -u all

## lint: runs a number of code quality checks against the source code
lint: $(GOBIN)/golangci-lint
	golangci-lint cache clean
	golangci-lint run

$(GOBIN)/golangci-lint:
	cd && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

## loc: displays the lines of code (LoC) count
loc:
	@loc --exclude _sample_configs/ _site/ docs/ Makefile *.md
# cargo install --locked loc

## install: installs a local version of the app
install:
	$(eval GOVERS = $(shell go version))
	@echo "$$HEADER"
	@echo "Installing ${APP} with ${GOVERS}..."
	@go clean
	@go install -ldflags="-s -w"
	@mv $(GOBIN)/${APP} $(GOBIN)/${APP}
	$(eval INSTALLPATH = $(shell which ${APP}))
	@echo "${APP} installed into ${INSTALLPATH}"

## uninstall: uninstals a locally-installed version
uninstall:
	@rm $(GOBIN)/${APP}

## docker-build: builds in docker
docker-build:
	@echo "Building ${APP} in Docker..."
	docker build -t ngtools:build --build-arg=version=master -f Dockerfile.build .
	@echo "Done with docker build"

## docker-install: installs a local version of the app from docker build
docker-install:
	@echo "Installing..."
	docker create --name ngtools_build ngtools:build
	docker cp ngtools_build:/usr/local/bin/ngtools ~/.local/bin/
	$(eval INSTALLPATH = $(shell which ${APP}))
	@echo "${APP} installed into ${INSTALLPATH}"
	docker rm ngtools_build

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

