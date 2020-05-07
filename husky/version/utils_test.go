package version

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetRepoName(t *testing.T) {
	NewWithT(t).Expect(GetRepoName("git@github.com:go-courier/husky.git")).To(Equal("go-courier/husky"))
	NewWithT(t).Expect(GetRepoName("https://github.com/go-courier/husky.git")).To(Equal("go-courier/husky"))
}
