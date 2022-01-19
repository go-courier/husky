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

dep:
	go get -u ./...
	go get -t ./...