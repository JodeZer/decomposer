package helper

import (
	"io/ioutil"
	"strings"

	"github.com/JodeZer/bag/lib/bag"
)

type DirExplorer struct {
	// prefix at mean time
	srcDir string
	// without prefix
	dirSlice  *bag.StringSlice
	fileSlice *bag.StringSlice
}

func fullPathComplementer(srcDir string) bag.StringValMapper {
	return func(str string) string {
		return srcDir + "/" + str
	}
}

func unixHidedDir() bag.StringValFilter {
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
		dirSlice:  bag.MakeStringSlice(0, 1),
		fileSlice: bag.MakeStringSlice(0, 1),
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

func (d *DirExplorer) GetFullPathDirs(filter bag.StringValFilter) *bag.StringSlice {
	return d.dirSlice.MapVal(fullPathComplementer(d.srcDir)).FilterVal(filter)
}

func (d *DirExplorer) GetFullPathFiles(filter bag.StringValFilter) *bag.StringSlice {
	return d.fileSlice.MapVal(fullPathComplementer(d.srcDir)).FilterVal(filter)
}
