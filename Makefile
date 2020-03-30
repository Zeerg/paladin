.PHONY: build build-alpine clean test help default

BIN_NAME=paladin

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "paladin/paladin"

default: clean build-art build

help:
	@echo 'Management commands for paladin:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make build-art       Build atomic red team static fs.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

clean:
	rm -rf art/
	rm -rf statik/
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

build-art:
	@echo "building Atomic Red Team Static FS"
	git clone https://github.com/redcanaryco/atomic-red-team.git art
	~/go/bin/statik -include=*.yaml -src=art/atomics
build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/Zeerg/paladin/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/Zeerg/paladin/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure

test:
	go test ./...

