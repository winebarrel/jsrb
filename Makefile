SHELL   := /bin/bash
VERSION := v1.1.0
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" ./cmd/jsrt

.PHONY: package
package: clean build
	gzip jsrt -c > jsrt_$(VERSION)_$(GOOS)_$(GOARCH).gz
	sha1sum jsrt_$(VERSION)_$(GOOS)_$(GOARCH).gz > jsrt_$(VERSION)_$(GOOS)_$(GOARCH).gz.sha1sum

.PHONY: clean
clean:
	rm -f jsrt
