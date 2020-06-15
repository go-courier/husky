VERSION=$(shell cat .version)

install:
	go install -v -ldflags "-X github.com/go-courier/husky/version.Version=${VERSION}"

lint:
	husky hook pre-commit
	husky hook commit-msg

test:
	go test -v -race ./...

cover:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

release:
	git push
	git push origin v${VERSION}