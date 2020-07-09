PKG=$(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION=v$(shell cat .version)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -ldflags "-X ${PKG}/version.Version=${VERSION}"

MAIN_ROOT ?= ./cmd/husky


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(MAKE) build.husky
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(MAKE) build.husky
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(MAKE) build.husky

	cd ./cmd/husky/out && pwd

build.husky:
	cd $(MAIN_ROOT) && $(GOBUILD) -o ./out/husky-$(GOOS)-$(GOARCH) 

install: build
	mv $(MAIN_ROOT)/husky ${GOPATH}/bin/husky

deps:
	cd $(MAIN_ROOT) && go get -u

release:
	git push
	git push origin ${VERSION}