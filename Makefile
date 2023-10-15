include scripts/*

BINARY_NAME=go-migrate
DESTINATION=./bin/local/${BINARY_NAME}
VERSION=$(shell git describe --always --tags | sed 's/-/+/')
LINKER_FLAGS=-X github.com/prastuvwxyz/go-migrate/internal/cli.Version=${VERSION}

all: tidy build

tidy: pull-depedencies

build: compile-local

compile-local:
	CGO_ENABLED=0 go build -ldflags "${LINKER_FLAGS}" -o ${DESTINATION} ./cmd/go-migrate
