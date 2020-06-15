package conventionalcommit

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseCommitMsg(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat(parser-test): add ability to parse arrays`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.Scope).To(Equal("parser-test"))
		NewWithT(t).Expect(cm.Header).To(Equal("add ability to parse arrays"))
	})

	t.Run("BREAKING CHANGE", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat: allow provided config object to extend other configs

BREAKING CHANGE: extends key in config file is now used for extending other config files`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.BreakingChange).To(BeTrue())
	})

	t.Run("BREAKING CHANGE", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat!: allow provided config object to extend other configs`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.BreakingChange).To(BeTrue())
	})
}
