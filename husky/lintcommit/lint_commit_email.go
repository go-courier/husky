package lintcommit

import (
	"fmt"
	"regexp"
)

func CreateLintCommitEmail(regexpStr string) func(email string) error {
	re := regexp.MustCompile(regexpStr)

	return func(email string) error {
		if re.MatchString(email) {
			return nil
		}
		return fmt.Errorf("invalid email `%s`, domain should be match `%s`, please use `git config user.email <email>` to fix", email, regexpStr)
	}
}
