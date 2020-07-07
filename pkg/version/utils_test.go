package version

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetRepoName(t *testing.T) {
	NewWithT(t).Expect(getBaseURI("git@github.com:go-courier/husky.git")).To(Equal("https://github.com/go-courier/husky"))
	NewWithT(t).Expect(getBaseURI("https://github.com/go-courier/husky.git")).To(Equal("https://github.com/go-courier/husky"))
}
