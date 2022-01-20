# Husky

[![GoDoc Widget](https://godoc.org/github.com/go-courier/husky?status.svg)](https://godoc.org/github.com/go-courier/husky)
[![codecov](https://codecov.io/gh/go-courier/husky/branch/master/graph/badge.svg)](https://codecov.io/gh/go-courier/husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-courier/husky)](https://goreportcard.com/report/github.com/go-courier/husky)

Husky.js like, but pure in golang

## Usage

install

```
go install github.com/go-courier/husky/cmd/husky
```

install git hooks

```
husky init
```

## Configuration `.husky.toml`

`.husky.yaml` supported too.

```toml
# version-file which will write or read current semver
version-file = "internal/version/version"

# hook scripts
[hooks]

# after version calc,
# with use the {{ .Version }} to upgrade other files.
post-version = [
    "sed -i -e 's/\"version\": \"[^\"]*\"/\"version\": \"{{ .Version }}\"/g' testdata/pacakge.json",
    "sed -i -e 's/version: [^\\n]*/version: {{ .Version }}/g' testdata/pubspec.yaml"
]

# git hook pre commit
pre-commit = [
    "golangci-lint run",
    "husky lint-staged",
]

# git hook commit msg
commit-msg = [
    "husky lint-commit",
]

# list staged files do some pre-process and git add
[lint-staged]
"*.go" = [
    "goimports -l -w",
    "gofmt -l -w",
]

# commit msg rule default support conventional commits
[lint-commit]
# could check if this exists
# email = "^(.+@gmail.com|.+@qq.com)$"
# optional custom types check regex
# types = "^(feat|fix|build|chore|ci|docs|perf|refactor|revert|style|test)$"
# optional header check regex
# header = "^(?P<type>\w+)(\((?P<scope>[\w/.-]+)\))?(?P<breaking>!)?:( +)?(?P<header>.+)"
```

Commit msg rule follow <https://www.conventionalcommits.org/en/v1.0.0/>

```
type(scope?): header

body?

footer?
```
