package version

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestListCommit(t *testing.T) {
	lastVersion, _, err := LastVersion()
	if err != nil {
		return
	}
	spew.Dump(lastVersion)
	spew.Dump(ListCommit("v" + lastVersion.String()))
}
