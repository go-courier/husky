package lintcommit

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"regexp"

	"github.com/go-courier/husky/pkg/conventionalcommit"
	"github.com/go-courier/husky/pkg/log"
)

type LintCommit struct {
	Email  string `yaml:"email,omitempty"`
	Types  string `yaml:"types,omitempty"`
	Header string `yaml:"header,omitempty"`
}

func (lintCommit LintCommit) NewLint(ctx context.Context) func() error {
	var lintCommitEmail func(s string) error

	logger := log.LoggerFromContext(ctx).WithName("LintStaged")

	if lintCommit.Email != "" {
		lintCommitEmail = CreateLintCommitEmail(lintCommit.Email)
	}
	if lintCommit.Header != "" {
		conventionalcommit.HeaderRegex = regexp.MustCompile(lintCommit.Header)
	}
	if lintCommit.Types != "" {
		conventionalcommit.TypesRegex = regexp.MustCompile(lintCommit.Types)
	}

	return func() error {
		if lintCommitEmail != nil {
			email, err := getGitUserEmail()
			if err != nil {
				logger.Error(err, "failed to get git user email")
				return err
			}
			if err := lintCommitEmail(email); err != nil {
				return err
			}
		}
		commitMsg, err := getGitLastCommitMsg()
		if err != nil {
			return err
		}
		return LintCommitMsg(commitMsg)
	}
}

func getGitUserEmail() (string, error) {
	ret, err := exec.Command("git", "config", "--get", "user.email").CombinedOutput()
	return string(bytes.TrimSpace(ret)), err
}

func getGitLastCommitMsg() (string, error) {
	ret, err := os.ReadFile(".git/COMMIT_EDITMSG")
	return string(ret), err
}
