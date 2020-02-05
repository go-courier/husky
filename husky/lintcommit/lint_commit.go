package lintcommit

import (
	"bytes"
	"io/ioutil"
	"os/exec"
)

type LintCommit struct {
	Email string `yaml:"email,omitempty"`
}

func (lintCommit LintCommit) NewLint() func() error {
	var lintCommitEmail func(s string) error

	if lintCommit.Email != "" {
		lintCommitEmail = CreateLintCommitEmail(lintCommit.Email)
	}

	return func() error {
		if lintCommitEmail != nil {
			email, err := getGitUserEmail()
			if err != nil {
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
	ret, err := ioutil.ReadFile(".git/COMMIT_EDITMSG")
	return string(ret), err
}
