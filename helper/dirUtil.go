package helper

import (
	"io/ioutil"
	"strings"

	"github.com/JodeZer/decomposer/helper/helperF"
)

type DirExplorer struct {
	// prefix at mean time
	srcDir string
	// without prefix
	dirSlice  *StringSlice
	fileSlice *StringSlice
}

func fullPathComplementer(srcDir string) helperF.StringMender {
	return func(str string) string {
		return srcDir + "/" + str
	}
}

func unixHidedDir() helperF.StringFilter {
	return func(str string) bool {
		if str[0] == '.' {
			return true
		}
		sps := strings.Split(str, "/")
		return sps[len(sps)-1][0] == '.'
	}
}

func MustConstructDirExplorer(dir string) *DirExplorer {
	explorer := &DirExplorer{
		srcDir:    dir,
		dirSlice:  MakeStringSlice(0, 1),
		fileSlice: MakeStringSlice(0, 1),
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		if fi.IsDir() {
			explorer.dirSlice.Append(fi.Name())
		} else {
			explorer.fileSlice.Append(fi.Name())
		}

	}
	return explorer
}

func (d *DirExplorer) GetFullPathDirs(filter helperF.StringFilter) *StringSlice {
	return d.dirSlice.GetMendedSlice(fullPathComplementer(d.srcDir)).GetFilteredSlice(filter)
}

func (d *DirExplorer) GetFullPathFiles(filter helperF.StringFilter) *StringSlice {
	return d.fileSlice.GetMendedSlice(fullPathComplementer(d.srcDir)).GetFilteredSlice(filter)
}
