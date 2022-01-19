GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

build:
	goreleaser build --snapshot --rm-dist

test:
	go test -v -race ./...

cover:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...

install: build
	mv ./bin/husky_$(GOOS)_$(GOARCH)/husky ${GOPATH}/bin/husky

test.version: install
	husky --verbose=10 version --skip-pull --skip-commit

dep:
	go get -u ./...
	go get -t ./...