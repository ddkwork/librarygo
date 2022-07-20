package cyclicimport

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

type (
	Interface interface {
		LogError(body string) (ok bool)
		WriteAppend(filePath, buf string) (ok bool)
		CheckError(err error) bool
		CheckError2(_ any, err error) bool
		IsAndroid() bool
		GetTimeNowString() string
	}
	object struct{}
)

func New() Interface {
	return &object{}
}

func (o *object) GetTimeNowString() string {
	return time.Now().Format("2006-01-02 15:04:05 ")
}
func (o *object) IsAndroid() bool {
	return runtime.GOOS == "android"
}

func (o *object) LogError(body string) (ok bool) {
	if o.IsAndroid() {
		return
	}
	ColorBody := fmt.Sprintf("\x1b[91m%s\x1b[0m", body) //red
	fmt.Println(ColorBody)
	return o.WriteAppend("log.log", body)
}

func (o *object) WriteAppend(filePath, buf string) (ok bool) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func() {
		if file == nil {
			o.CheckError(errors.New("file == nil "))
			return
		}
		o.CheckError(file.Close())
	}()
	if !o.CheckError(err) {
		println("maybe dir is not created " + err.Error())
		if !o.CheckError(os.MkdirAll(filePath, os.ModePerm)) {
			return
		}
	}
	return o.CheckError2(file.WriteString(buf))
}

func (o *object) CheckError2(_ any, err error) bool { return o.CheckError(err) }

func (o *object) CheckError(err error) bool {
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return false
	}
	return true
}
