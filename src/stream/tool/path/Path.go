package path

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"os"
	"path/filepath"
	"strings"
)

type (
	Interface interface {
		CreatDirectory(path string) bool           //创建目录
		CurrentDirectory() string                  //返回当前目录
		GetCurrentDirectory() (ok bool)            //获取当前目录
		GetCurrentRunningPath() (ok bool)          //获取当前运行路径
		CurrentRunningPath() string                //返回当前运行路径
		GetFilepath(fileName string) (ok bool)     //获取文件路径
		Filepath() string                          //返回文件路径
		GetFileDirectory(Filepath string) string   //返回文件目录
		GetFilename(Filepath string) string        //返回文件名
		GetExtensionPath(Filepath string) string   //返回文件扩展名
		GetAbsolutePath(Filepath string) (ok bool) //获取绝对路径
		AbsolutePath() string                      //返回绝对路径
		GetUNCPath(fileName string) (ok bool)      //获取UNC路径
		UNCPath() string                           //返回UNC路径
	}
	object struct {
		currentDirectory   string
		absolutePath       string
		currentRunningPath string
		filepath           string
		uncPath            string
		err                error
	}
)

func New() Interface { return &object{} }

func (o *object) UNCPath() string { return o.uncPath }

func (o *object) GetUNCPath(fileName string) (ok bool) {
	if !o.GetCurrentRunningPath() {
		return
	}
	if !o.GetFilepath(fileName) {
		return
	}
	o.uncPath = strings.Replace(o.filepath, `\`, `\\`, -1)
	return true
}

func (o *object) CreatDirectory(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.MkdirAll(filepath.Dir(path), os.ModePerm) == nil {
			return mycheck.Error2(os.Create(path))
		}
	}
	return true
}
func (o *object) CurrentDirectory() string   { return o.currentDirectory }
func (o *object) AbsolutePath() string       { return o.absolutePath }
func (o *object) CurrentRunningPath() string { return o.currentRunningPath }

func (o *object) GetCurrentRunningPath() (ok bool) {
	dir, err := os.Executable()
	if !mycheck.Error2(dir, err) {
		return
	}
	o.currentRunningPath = filepath.Dir(dir)
	return true
}

func (o *object) GetCurrentDirectory() (ok bool) {
	o.currentDirectory, o.err = os.Getwd()
	return mycheck.Error(o.err)
}

func (o *object) Filepath() string {
	return o.filepath
}
func (o *object) GetFilepath(fileName string) (ok bool) {
	if !o.GetCurrentDirectory() {
		return
	}
	o.filepath = filepath.Join(o.currentDirectory, fileName)
	return true
}
func (o *object) GetFileDirectory(Filepath string) string {
	return filepath.Dir(Filepath)
}
func (o *object) GetFilename(Filepath string) string {
	return filepath.Base(Filepath)
}
func (o *object) GetAbsolutePath(Filepath string) (ok bool) {
	o.absolutePath, o.err = filepath.Abs(Filepath)
	return mycheck.Error(o.err)
}
func (o *object) GetExtensionPath(Filepath string) string {
	return filepath.Ext(Filepath)
}
