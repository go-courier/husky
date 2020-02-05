lint: hook
	export HUSKY_DEBUG=1
	husky run pre-commit
	husky run commit-msg

hook: install
	husky init

install:
	go install github.com/go-courier/husky

test:
	go test -v -race ./...

cover:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
