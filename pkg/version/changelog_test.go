package version_test

import (
	"bytes"
	"testing"

	"github.com/go-courier/husky/pkg/version"
)

func TestUpdateChangeLog(t *testing.T) {
	commits, _ := version.ListCommit("")

	f := bytes.NewBuffer(nil)

	t.Run("from empty", func(t *testing.T) {
		nextVer, sections := version.CalcNextVer(commits, nil)
		t.Log(nextVer)

		_ = version.UpdateChangeLog(f, nextVer, nil, sections)
		t.Log(f.String())
	})

	t.Run("from exists", func(t *testing.T) {
		lastVer, list, _ := version.ResolveVersionAndCommits()
		nextVer, sections := version.CalcNextVer(list, lastVer)
		t.Log(nextVer)

		_ = version.UpdateChangeLog(f, nextVer, lastVer, sections)
		t.Log(f.String())
	})
}
