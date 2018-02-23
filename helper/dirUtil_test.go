package helper

import "testing"

func TestDir(t *testing.T) {
	de := MustConstructDirExplorer("/Users/ezbuy/Projects/ezbuy/goflow/src/github.com/JodeZer/decomposer")
	t.Logf("%+v", de.GetFullPathDirs(unixHidedDir()).GetRaw())
	t.Logf("%+v", de.GetFullPathFiles(nil).GetRaw())
}
