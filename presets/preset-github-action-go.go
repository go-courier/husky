package presets

func init() {
	Register("github-action-go", Preset{
		".github/workflows/ci.yml": []byte(`
name: test

on: push

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: '^1.14.0'
      - run: go install github.com/go-courier/husky
      - run: make cover
      - uses: codecov/codecov-action@v1
        with:
          file: ./coverage.txt
          fail_ci_if_error: true
`),
	})
}
