package helper

import (
	"io/ioutil"
	"strings"

	"github.com/JodeZer/decomposer/helper/helperF"
)

type FileExplorer struct {
	// explore source dir
	srcDir string
	// ignore file suffix
	ignoreU UniqueMap
	// target file suffix
	targetU UniqueMap
}

func NewFileExplorer(srcDir string) *FileExplorer {
	return &FileExplorer{
		srcDir:  srcDir,
		ignoreU: NewUniqueMap(1),
		targetU: NewUniqueMap(1),
	}
}

func (this *FileExplorer) SetIgnore(suffix ...string) {
	this.ignoreU.MAdd(helperF.ConvertStringInterface(suffix...))
}

func (this *FileExplorer) SetTarget(suffix ...string) {
	this.targetU.MAdd(helperF.ConvertStringInterface(suffix...))
}

func (this *FileExplorer) Explore() []string {
	return nil
}

func (this *FileExplorer) MustListDirWithDepth(srcDir string) []string {
	fileInfos, err := ioutil.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}
	res := make([]string, 0, 5)
	for _, fi := range fileInfos {
		if fi.IsDir() {
			res = append(res, this.MustListDirWithDepth(srcDir+"/"+fi.Name())...)
		}
	}
	return res
}

// unexported method
// ===========================

func (this *FileExplorer) mustFindLegalFile(srcDir string) []string {
	return nil
}

func (this *FileExplorer) ignore(file string) bool {
	for k, _ := range this.ignoreU {
		if strings.HasSuffix(file, k.(string)) {
			return true
		}
	}
	return false
}

func (this *FileExplorer) target(file string) bool {
	for k, _ := range this.targetU {
		if strings.HasSuffix(file, k.(string)) {
			return true
		}
	}
	return false
}

func (this *FileExplorer) findAll() []string {
	return nil
}
