APPLICATION = ipam
ORGANIZATION = RackHD

TTY = $(shell if [ -t 0 ]; then echo "-ti"; fi)

DOCKER_DIR = /go/src/github.com/${ORGANIZATION}/${APPLICATION}
DOCKER_IMAGE = rackhd/golang:1.7.0-wheezy
DOCKER_CMD = docker run ${TTY} --rm -v ${PWD}:${DOCKER_DIR} -w ${DOCKER_DIR} ${DOCKER_IMAGE}

.PHONY: shell deps deps-local build build-local lint lint-local test test-local release

default: deps test build

shell:
	@docker run --rm -ti -v ${PWD}:${DOCKER_DIR} -w ${DOCKER_DIR} ${DOCKER_IMAGE} /bin/bash

deps:
	@${DOCKER_CMD} make deps-local

deps-local:
	@if ! [ -f glide.yaml ]; then glide init --non-interactive; fi
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
	@ginkgo -race -trace -randomizeAllSpecs -r -cover

coveralls:
	@go get github.com/mattn/goveralls
	@go get github.com/modocache/gover
	@go get golang.org/x/tools/cmd/cover
	@gover
	@goveralls -coverprofile=gover.coverprofile -service=travis-ci

release: build
	@docker build -t rackhd/${APPLICATION} .

run: release
	@docker-compose up

mongo:
	@docker exec -it mongodb mongo ipam
