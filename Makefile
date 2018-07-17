GOOS ?= $(shell go env GOOS)
DEFAULT_ARCH := amd64
BUILD_TIME = $(shell date -u +%Y%m%d.%H%M%S)
VERSION = $(shell cat version)
LDFLAGS = -ldflags "-X main.buildTime=$(BUILD_TIME) -X main.version=$(VERSION)"
BINARY = scmt-$(GOOS)-$(VERSION)

default: build

clean:
	rm -rf build/*

.PHONY: build
build: 
	GOOS=${GOOS} GOARCH=${DEFAULT_ARCH} go build ${LDFLAGS} -o build/${BINARY} .
