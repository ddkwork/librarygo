package file

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ddkwork/librarygo/src/check"
	mypath "github.com/ddkwork/librarygo/src/stream/tool/path"
	"github.com/hjson/hjson-go"
	"go/format"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type (
	Interface interface {
		WriteTruncate(name string, data any) (ok bool)
		WriteAppend(name string, data any) (ok bool)
		WriteGoCode(name string, data any) (ok bool)
		WriteBinary(name string, data any) (ok bool)
		WriteJson(name string, Obj any) (ok bool)
		WriteHjson(name string, Obj any) (ok bool)
		ToLines(data any) (lines []string, ok bool)
		GoCode() string
		Copy(source, destination string) (ok bool)
	}
	object struct{ goCode string }
)

func (o *object) Copy(source, destination string) (ok bool) {
	base := filepath.Base(source)
	return check.Error(filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		split := strings.Split(path, base)
		dst := filepath.Join(destination, base, split[1])
		switch {
		case info.IsDir():
			if !mypath.New().CreatDirectory(dst) {
				return err
			}
		default:
			buf, err := ioutil.ReadFile(path)
			if !check.Error(err) {
				return check.Object()
			}
			f, err := os.Create(dst)
			if !check.Error(err) {
				return err
			}
			if !check.Error2(f.Write(buf)) {
				return check.Object()
			}
			if !check.Error(f.Close()) {
				return check.Object()
			}
		}
		return err
	}))
}

func (o *object) buffer(data any) *bytes.Buffer {
	switch data.(type) {
	case string:
		return bytes.NewBufferString(data.(string))
	case []byte:
		return bytes.NewBuffer(data.([]byte))
	}
	return bytes.NewBufferString("error file data type " + fmt.Sprintf("%t", data))
}

func (o *object) WriteTruncate(name string, data any) (ok bool) {
	if !check.Error(os.Truncate(name, 0)) {
		return
	}
	return o.WriteAppend(name, data)
}
func (o *object) WriteAppend(name string, data any) (ok bool) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if !check.Error(err) {
		if !mypath.New().CreatDirectory(name) {
			return o.WriteAppend(name, data)
		}
	}
	if !check.Error2(f.Write(o.buffer(data).Bytes())) {
		return
	}
	if !check.Error2(f.WriteString("\n")) {
		return
	}
	return check.Error(f.Close())
}
func (o *object) WriteGoCode(name string, data any) (ok bool) {
	b, err := format.Source(o.buffer(data).Bytes())
	if !check.Error(err) {
		return
	}
	return o.WriteAppend(name, b)
}
func (o *object) WriteBinary(name string, data any) (ok bool) {
	file, err := os.Create(name)
	if !check.Error(err) {
		return
	}
	if !check.Error2(file.Write(o.buffer(data).Bytes())) {
		return
	}
	return check.Error(file.Close())
}
func (o *object) ToLines(data any) (lines []string, ok bool) {
	newReader := bufio.NewReader(o.buffer(data))
	for {
		line, _, err := newReader.ReadLine()
		switch err {
		case io.EOF:
			return lines, true
		default:
			if !check.Error(err) {
				return
			}
		}
		lines = append(lines, string(line))
	}
}
func (o *object) WriteJson(name string, Obj any) (ok bool) {
	var oo interface{}
	switch reflect.TypeOf(Obj).Kind() {
	case reflect.Struct:
		oo = &Obj
	case reflect.Ptr:
		oo = Obj
	}
	data, err := json.MarshalIndent(oo, " ", " ")
	if !check.Error(err) {
		return
	}
	return o.WriteAppend(name, data)
}
func (o *object) WriteHjson(name string, Obj any) (ok bool) {
	data, err := hjson.MarshalWithOptions(Obj, hjson.EncoderOptions{
		Eol:            "",
		BracesSameLine: false,
		EmitRootBraces: false,
		QuoteAlways:    false,
		IndentBy:       " ",
		AllowMinusZero: false,
		UnknownAsNull:  false,
	})
	if !check.Error(err) {
		return
	}
	return o.WriteAppend(name, data)
}
func (o *object) GoCode() string { return o.goCode }
func New() Interface {
	return &object{
		goCode: "",
	}
}
