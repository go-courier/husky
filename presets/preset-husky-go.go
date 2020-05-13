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
  release:
    - git push
    - git push origin $(git describe --tags --abbrev=0)

hooks:
  pre-commit:
    - golangci-lint run
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
