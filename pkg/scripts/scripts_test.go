package scripts

import (
	"os"
	"testing"

	"github.com/onsi/gomega"
)

func TestRunScript(t *testing.T) {
	t.Run("simple echo", func(t *testing.T) {
		os.Setenv("TEST", "test")

		err := RunScript(`echo "${TEST}"`)
		gomega.NewWithT(t).Expect(err).To(gomega.BeNil())
	})

	t.Run("ls", func(t *testing.T) {
		err := RunScript(`ls`)
		gomega.NewWithT(t).Expect(err).To(gomega.BeNil())
	})

	t.Run("xxvasdad", func(t *testing.T) {
		err := RunScript(`__xasdasd`)
		gomega.NewWithT(t).Expect(err).NotTo(gomega.BeNil())
	})
}
