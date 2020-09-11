package scripts

import (
	"context"
	"os"
	"testing"

	"github.com/onsi/gomega"
)

func TestRunScript(t *testing.T) {
	ctx := context.Background()

	t.Run("simple echo", func(t *testing.T) {
		os.Setenv("TEST", "test")

		err := RunScript(ctx, `echo "${TEST}"`)
		gomega.NewWithT(t).Expect(err).To(gomega.BeNil())
	})

	t.Run("ls", func(t *testing.T) {
		err := RunScript(ctx, `ls`)
		gomega.NewWithT(t).Expect(err).To(gomega.BeNil())
	})

	t.Run("xxvasdad", func(t *testing.T) {
		err := RunScript(ctx, `__xasdasd`)
		gomega.NewWithT(t).Expect(err).NotTo(gomega.BeNil())
	})
}
