package CommentDeletor

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type (
	Interface interface {
		Delete(root string) (ok bool)
		DeleteKepSpace(root string) (ok bool)
	}
	skipInfo struct {
		index int
		code  string
	}
	object struct {
		body        string
		lines       []string
		path        string
		paths       []string
		isDebug     bool
		index       int
		skipLines   []skipInfo
		isKeepSpace bool
	}
)

func New() Interface {
	return newObject()
}
func newObject() *object {
	return &object{
		body:      "",
		lines:     nil,
		path:      "",
		paths:     make([]string, 0),
		isDebug:   false,
		index:     0,
		skipLines: make([]skipInfo, 0),
	}
}

func (o *object) DeleteKepSpace(root string) (ok bool) {
	o.isKeepSpace = true
	return o.Delete(root)
}
func (o *object) Paths() []string { return o.paths }

func main() {
	d := New()
	//d.Delete("include")
	//d.Delete("D:\\vt\\xa-tmp")
	//d.Delete("D:\\workspace\\workspace\\src\\cpp_work\\src\\vxk")
	//d.Delete("C:\\Users\\Admin\\Desktop\\HyperDbgDev")
	d.Delete("C:\\Users\\Admin\\Desktop\\nvme\\ColdStorageManager-main\\NVMeQuery.dll source (from CLion project)")
}
func (o *object) Delete(root string) (ok bool) {
	return mycheck.Error(filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".clang-format" && !o.isKeepSpace {
			if !mycheck.Error(os.Remove(path)) {
				return err
			}
		}
		o.path = path
		if !o.hasExt() {
			return err
		}
		o.paths = append(o.paths, path)
		if !o.CleanFile(path) {
			return err
		}
		return nil
	}))
}
func (o *object) CleanFile(path string) (ok bool) {
	if !o.FileToLines(path) {
		return
	}
	o.FindGroup()
	o.FindSingularLines()
	//o.debug()
	if !o.RemoveSpace() {
		return
	}
	//go func() {
	//if !o.ClangFormat() {
	//	return
	//}
	//}()
	paths := o.paths
	*o = *newObject() //todo change to reset
	o.paths = paths
	return true
}
func (o *object) FindGroup() {
	isBlock := false
	for i := 0; i < o.Size(); i++ {
		line := o.lines[i]
		_, start, end := o.startEnd(line) //不去空格会导致判断失误，skipLines重复填充过大，ResetLine的时候把不该清零的干掉
		if start == "/**" && strings.LastIndex(line, "*/") < 0 {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
			isBlock = true
			continue
		}
		if end == "*/" && isBlock {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
			o.ResetLine()
			isBlock = false
		}
		if isBlock {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
		}
	}

	for i := 0; i < o.Size(); i++ {
		line := o.lines[i]
		_, start, end := o.startEnd(line)
		if start == "/*" && strings.LastIndex(line, "*/") < 0 {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
			isBlock = true
			continue
		}
		if end == "*/" && isBlock {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
			o.ResetLine()
			isBlock = false
		}
		if isBlock {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
		}
	}
}

func (o *object) FindSingularLines() {
	for i := 0; i < o.Size(); i++ {
		line := o.lines[i]
		_, start, _ := o.startEnd(line)
		if start == "//" {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
		}
	}
	o.ResetLine()

	for i := 0; i < o.Size(); i++ {
		line := o.lines[i]
		_, start, end := o.startEnd(line)
		if start == "/*" && end == "*/" {
			o.skipLines = append(o.skipLines, skipInfo{index: i, code: line})
		}
	}
	o.ResetLine()
}
func (o *object) RemoveSpace() (ok bool) {
	file, err := os.Create(o.path)
	if !mycheck.Error(err) {
		return
	}
	for _, line := range o.lines {
		if line != "" && !o.isKeepSpace {
			if !mycheck.Error2(file.WriteString(line)) {
				return
			}
			if !mycheck.Error2(file.WriteString("\n")) {
				return
			}
			if line == "}" {
				if !mycheck.Error2(file.WriteString("\n")) {
					return
				}
			}
		}
	}
	return mycheck.Error(file.Close())
}
