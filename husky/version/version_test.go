package version

import "testing"

func TestVersion(t *testing.T) {
	t.Log(Version(VersionOpt{}))
}
