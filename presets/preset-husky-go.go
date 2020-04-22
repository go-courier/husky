package presets

func init() {
	Register("husky-go", Preset{
		".husky.yaml": []byte(`
scripts:
  lint:
    - husky hook pre-commit
    - husky hook commit-msg
  test:
    - go test -v -race ./...
  cover:
    - go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
  install:
    - go install -v

hooks:
  pre-commit:
    - go vet ./...
    - husky lint-staged
  commit-msg:
    - husky lint-commit

lint-staged:
  "*.go":
    - gofmt -l -w

lint-commit:
  email: "^(.+@gmail.com|.+@qq.com)$"
`),
	})
}
