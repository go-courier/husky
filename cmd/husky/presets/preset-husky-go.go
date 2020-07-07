package presets

func init() {
	Register("husky-go", Preset{
		".husky.yaml": []byte(`
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
