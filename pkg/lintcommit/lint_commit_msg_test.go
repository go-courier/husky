package lintcommit

import (
	"regexp"
	"testing"

	"github.com/go-courier/husky/pkg/conventionalcommit"
	. "github.com/onsi/gomega"
)

func TestCheckCommitMsg(t *testing.T) {
	t.Run("invalid format", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg(": test")).NotTo(BeNil())
	})

	t.Run("invalid type", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("hehe: test")).NotTo(BeNil())
	})

	t.Run("valid header", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("chore: test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("fix(account): #IEP-111 test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("feat(account)!: #IEP-111 test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("fix(test/account): #IEP-111 test")).To(BeNil())
	})
}

func TestCustomCheckCommitMsg(t *testing.T) {
	conventionalcommit.TypesRegex = regexp.MustCompile(`^(custom|types)$`)
	conventionalcommit.HeaderRegex = regexp.MustCompile(`^(?P<issue>[A-Z]+-[\d]+)( +)(?P<type>\w+)(\((?P<scope>[\w/.-]+)\))?(?P<breaking>!)?:( +)?(?P<header>.+)`)

	t.Run("invalid format", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg(": test")).NotTo(BeNil())
		NewWithT(t).Expect(LintCommitMsg("custom: test")).NotTo(BeNil())
	})

	t.Run("invalid type", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("IEP-1 feat: test")).NotTo(BeNil())
	})

	t.Run("valid header", func(t *testing.T) {
		NewWithT(t).Expect(LintCommitMsg("IEP-1 custom: test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("IEP-11 custom(account): test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("IEP-111 types(account)!: test")).To(BeNil())
		NewWithT(t).Expect(LintCommitMsg("IEP-111 types(test/account): test")).To(BeNil())
	})
}
