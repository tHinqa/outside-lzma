package lzma

import "testing"

func TestInit(t *testing.T) {
	t.Logf("liblzma version: VersionString() %s; VersionNumber() %d", VersionString(), VersionNumber())
}
