APPLICATION = ipam
ORGANIZATION = RackHD

DOCKER_DIR = /go/src/github.com/${ORGANIZATION}/${APPLICATION}
DOCKER_IMAGE = skunkworxs/golang:1.7-wheezy
DOCKER_CMD = docker run -ti --rm -v ${PWD}:${DOCKER_DIR} -w ${DOCKER_DIR} ${DOCKER_IMAGE}

.PHONY: shell deps deps-local build build-local lint lint-local test test-local release

default: deps test build

shell:
	@docker run --rm -ti -v ${PWD}:${DOCKER_DIR} -w ${DOCKER_DIR} ${DOCKER_IMAGE} /bin/bash

deps:
	@${DOCKER_CMD} make deps-local

deps-local:
	@if ! [ -f glide.lock ]; then glide init --non-interactive; fi
	@glide install --strip-vcs --strip-vendor

build:
	@${DOCKER_CMD} make build-local

build-local:
	@go build -o bin/${APPLICATION} main.go

lint:
	@${DOCKER_CMD} make lint-local

lint-local:
	@gometalinter --vendor --fast --disable=dupl --disable=gotype --skip=grpc ./...

test:
	@${DOCKER_CMD} make test-local

test-local: lint-local
	@ginkgo -race -trace -randomizeAllSpecs -r

release: build
	@docker build -t skunkworxs/${APPLICATION} .

run: release
	@docker-compose up
