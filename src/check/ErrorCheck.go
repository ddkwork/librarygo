package check

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/ddkwork/librarygo/src/check/internal/cyclicimport"
	"github.com/ddkwork/librarygo/src/check/internal/table"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func Assert(t *testing.T) *assert.Assertions         { return assert.New(t) }
func Error(err interface{}) bool                     { return Default.Error(err) }
func Bool2(retCtx interface{}, ok bool) bool         { return Default.Bool2(retCtx, ok) }
func Error2(retCtx interface{}, err error) (ok bool) { return Default.Error2(retCtx, err) }
func Body() string                                   { return Default.Body() }
func List() []_ErrorList                             { return Default.List() }
func Object() error                                  { return Default.Object() }

type (
	Interface interface {
		Error(err interface{}) bool
		Bool2(retCtx interface{}, ok bool) bool
		Error2(retCtx interface{}, err error) (ok bool)
		Body() string
		List() []_ErrorList //todo filter list
		Object() error
	}
	_ErrorList struct {
		reason string
		body   string
	}
	object struct {
		body         string
		cyclicImport cyclicimport.Interface
		errorList    []_ErrorList
	}
)

func (o *object) Object() error { return errors.New(o.body) }
func (o *object) SetErrorList(errorListInfo _ErrorList) {
	o.errorList = append(o.errorList, errorListInfo)
}
func (o *object) List() []_ErrorList { return o.errorList }

var Default = New()

func New() Interface {
	return &object{
		body:         "",
		cyclicImport: cyclicimport.New(),
		errorList:    nil,
	}
}

func (o *object) Body() string { return o.body }
func (o *object) Error(err interface{}) bool {
	if err == nil {
		return true
	}
	return o.setErrorInfo(err)
}

func (o *object) Error2(retCtx interface{}, err error) (ok bool) {
	if err == nil {
		return o.checkArg(retCtx)
	}
	return o.setErrorInfo(err) //干掉层层返回，一出错就能卡到代码了
}

func (o *object) Bool2(retCtx interface{}, ok bool) bool {
	if ok {
		return o.checkArg(retCtx)
	}
	return ok
}

func (o *object) checkArg(retCtx interface{}) bool {
	switch retCtx.(type) {
	case string:
		if retCtx == "" {
			return o.setErrorInfo("nil string")
		}
		if retCtx == "undefined" {
			return o.setErrorInfo("JsRun return undefined")
		}
		if retCtx == "{}" {
			return o.setErrorInfo("json Structure member names must be uppercase")
		}
	case int: //加入其它的整形？再观察下是否需要断言或者更多的判断？反射fun?
		if retCtx == 0 { //windows api似乎是返回0就是错误
			return o.setErrorInfo("Write 0 Bytes to file")
		}
	case []byte: //这些应该堆栈回溯判断下函数名？
		if retCtx.([]byte) == nil {
			return o.setErrorInfo("The network request did not return content")
			//个别爬虫网站确实需要忽略某个包，那么这里要返回true,也可以修改业务代码配合这个错误检查逻辑使之通用
		}
	case *template.Template:
		if retCtx.(*template.Template) == nil {
			return o.setErrorInfo("The html template file returns a null pointer, please object the content of the html file")
		}
	default:
		//检查别的类型
	}
	return true
}
func (o *object) FileToLines(src interface{}) (lines []string, ok bool) {
	NewSrc := make([]byte, 0)
	switch src.(type) {
	case string:
		NewSrc = []byte(src.(string))
	case []byte:
		NewSrc = src.([]byte)
	}
	reader := bytes.NewReader(NewSrc)
	newReader := bufio.NewReader(reader)
	for {
		line, _, err := newReader.ReadLine()
		switch err {
		case io.EOF:
			return lines, true
		default:
			if !o.Error(err) {
				return
			}
		}
		lines = append(lines, string(line))
	}
}
func (o *object) setErrorInfo(errorObject interface{}) (ok bool) {
	info := ""
	fnMakeLine := func(k, v string) string { return k + "\t" + v + "\n" }
	reason := ""
	switch errorObject.(type) {
	case error:
		reason = errorObject.(error).Error()
	case string:
		reason = errorObject.(string)
	}
	pc, _, _, ok := runtime.Caller(3)
	if !ok {
		return
	}
	fileLine := ""
	info += fnMakeLine("time", o.cyclicImport.GetTimeNowString())
	info += fnMakeLine("goroutine", fmt.Sprint(runtime.NumGoroutine()))
	info += fnMakeLine("", "               >>>>>>>>>>>>>>>>>>>>>>>>>>>> stack <<<<<<<<<<<<<<<<<<<<<<<<<<<")
	stack := strings.ReplaceAll(string(debug.Stack()), "\n\t", " ---> ")
	lines, ok := o.FileToLines(stack)
	if !ok {
		return
	}
	for _, s := range lines {
		if strings.Contains(s, runtime.FuncForPC(pc).Name()) {
			fileLine = s
		}
		info += "\t" + s + "\n"
	}
	info = fnMakeLine("reason", reason) + fnMakeLine("fileLine", fileLine) + info
	b := new(bytes.Buffer)
	table.Fprintf(b, table.BoxStyle, info,
		table.Centred,
		table.LeftJustified,
	)
	o.body = b.String()
	if !o.cyclicImport.LogError(b.String()) {
		return false
	}
	return false //false,con not be true
}
