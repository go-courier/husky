# Husky

[![GoDoc Widget](https://godoc.org/github.com/go-courier/husky?status.svg)](https://godoc.org/github.com/go-courier/husky)
[![Build Status](https://travis-ci.org/go-courier/husky.svg?branch=master)](https://travis-ci.org/go-courier/husky)
[![codecov](https://codecov.io/gh/go-courier/husky/branch/master/graph/badge.svg)](https://codecov.io/gh/go-courier/husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-courier/husky)](https://goreportcard.com/report/github.com/go-courier/husky)

Husky.js like, but pure in golang

## Usage

```
go install github.com/go-courier/husky
```


## Configuration `.husky.yaml`

```yaml
hooks:
  # hook scripts
  pre-commit:
    - husky lint staged
  commit-msg:
    - husky lint commit
  
# list staged files do some pre-process and git add
lint-staged:
  "*.go":
    - go fmt
    - go vet

# only support https://www.conventionalcommits.org/en/v1.0.0-beta.2/
lint-commit:
  # could check if this exists
  # regexp
  email: "^(.+@gmail.com|.+@qq.com)$"
```
