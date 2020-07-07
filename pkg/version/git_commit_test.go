package version

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestListCommit(t *testing.T) {
	lastVersion, _, _ := LastVersion()
	spew.Dump(lastVersion)
	spew.Dump(ListCommit("v" + lastVersion.String()))
}
