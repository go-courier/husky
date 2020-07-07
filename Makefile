PKG=$(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION=v$(shell cat .version)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -ldflags "-X ${PKG}/version.Version=${VERSION}"

MAIN_ROOT ?= ./cmd/husky

build:
	cd $(MAIN_ROOT) && $(GOBUILD)

install: build
	mv $(MAIN_ROOT)/husky ${GOPATH}/bin/husky

deps:
	cd $(MAIN_ROOT) && go get -u

release:
	git push
	git push origin ${VERSION}