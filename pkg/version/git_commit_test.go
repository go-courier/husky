package version

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestListCommit(t *testing.T) {
	lastVersion, _, _ := LastVersion()
	if lastVersion != nil {
		spew.Dump(lastVersion)
		spew.Dump(ListCommit("v" + lastVersion.String()))
	}
}
