package helper

import "io/ioutil"
import "github.com/JodeZer/decomposer/helper/helperF"

type DirExplorer struct {
	// prefix at mean time
	srcDir string
	// without prefix
	dir   []string
	files []string
}

func MustConstructDirExplorer(dir string) *DirExplorer {
	explorer := &DirExplorer{
		srcDir: dir,
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		if fi.IsDir() {
			explorer.dir = append(explorer.dir, fi.Name())
		}
		explorer.files = append(explorer.files, fi.Name())
	}
	return explorer
}

func (d *DirExplorer) GetFullPathDirs(filter helperF.StringFilter) []string {
	res := make([]string, 0, len(d.dir))
	for _, one := range d.dir {
		if !filter.Something()(one) {
			res = append(res, one)
		}
	}
	return res
}

func (d *DirExplorer) GetFullPathFiles(filter helperF.StringFilter) []string {
	return nil
}
