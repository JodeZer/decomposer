package helper

import (
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
	// dirUtil
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
	dirSlice := this.MustListDir(this.srcDir)
	dirSlice.Range(func(dir string) {

	})
}

func (this *FileExplorer) MustListDir(srcDir string) *StringSlice {
	de := MustConstructDirExplorer(srcDir)
	res := MakeStringSlice(0, 5)
	seedsDir := de.GetFullPathDirs(unixHidedDir())

	seedsDir.Range(func(str string) {
		res.AppendSlice(this.MustListDir(str))
	})
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
