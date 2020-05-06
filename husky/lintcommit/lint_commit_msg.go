package lintcommit

import (
	"github.com/go-courier/husky/husky/conventionalcommit"
)

func LintCommitMsg(commitMsg string) error {
	_, err := conventionalcommit.ParseCommit(commitMsg)
	if err != nil {
		return err
	}
	return nil
}
