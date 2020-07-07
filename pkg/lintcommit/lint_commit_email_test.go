package lintcommit

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCheckAuthorEmail(t *testing.T) {
	lint := CreateLintCommitEmail(`^*@gmail.com$`)

	NewWithT(t).Expect(lint("xxx@qq.com")).NotTo(BeNil())
	NewWithT(t).Expect(lint("xxx@gmail.com")).To(BeNil())
}
