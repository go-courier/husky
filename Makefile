PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git describe --always)-devel

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD = CGO_ENABLED=0 go build -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= husky

build:
	GOOS=linux GOARCH=amd64 $(MAKE) build.husky
	GOOS=linux GOARCH=arm64 $(MAKE) build.husky
	GOOS=darwin GOARCH=amd64 $(MAKE) build.husky

build.husky:
	$(GOBUILD) -o ./out/${WORKSPACE}/${WORKSPACE}-$(GOOS)-$(GOARCH) ./cmd/${WORKSPACE}

test:
	go test -v -race ./...

cover:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

install: build.husky
	mv ./out/${WORKSPACE}/${WORKSPACE}-$(GOOS)-$(GOARCH) ${GOPATH}/bin/husky

deps:
	go get -u ./...

release: install
	husky version --skip-tag
