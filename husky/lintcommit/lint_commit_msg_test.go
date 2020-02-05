package lintcommit

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCheckCommitMsg(t *testing.T) {
	t.Run("invalid format", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg(": test")).NotTo(BeNil())
	})

	t.Run("invalid type", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("hehe: test")).NotTo(BeNil())
	})

	t.Run("invalid rel", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("chore(deps): [IEP-xxx] test")).NotTo(BeNil())
	})

	t.Run("valid header", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("chore: test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("fix(account): [IEP-111] test")).To(BeNil())
	})
}
