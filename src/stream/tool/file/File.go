package file

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
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
		RaedToLines(path string) (lines []string, ok bool)
		GoCode() string
		Copy(source, destination string) (ok bool)
	}
	object struct{ goCode string }
)

func New() Interface {
	return &object{
		goCode: "",
	}
}
func (o *object) RaedToLines(path string) (lines []string, ok bool) {
	b, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	return o.ToLines(b)
}

func (o *object) Copy(source, destination string) (ok bool) {
	base := filepath.Base(source)
	return mycheck.Error(filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		split := strings.Split(path, base)
		dst := filepath.Join(destination, base, split[1])
		switch {
		case info.IsDir():
			if !mypath.New().CreatDirectory(dst) {
				return err
			}
		default:
			buf, err := ioutil.ReadFile(path)
			if !mycheck.Error(err) {
				return mycheck.Object()
			}
			f, err := os.Create(dst)
			if !mycheck.Error(err) {
				return err
			}
			if !mycheck.Error2(f.Write(buf)) {
				return mycheck.Object()
			}
			if !mycheck.Error(f.Close()) {
				return mycheck.Object()
			}
		}
		return err
	}))
}

func (o *object) buffer(data any) *bytes.Buffer { //todo replaced as stream pkg
	switch data.(type) {
	case string:
		return bytes.NewBufferString(data.(string))
	case []byte:
		return bytes.NewBuffer(data.([]byte))
	}
	return bytes.NewBufferString("error file data type " + fmt.Sprintf("%t", data))
}

func (o *object) WriteTruncate(name string, data any) (ok bool) {
	if err := os.Truncate(name, 0); err != nil {
		return o.WriteAppend(name, data)
	}
	return o.WriteAppend(name, data)
}
func (o *object) WriteAppend(name string, data any) (ok bool) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		if !mypath.New().CreatDirectory(name) {
			return
		}
		return o.WriteAppend(name, data)
	}
	if !mycheck.Error2(f.Write(o.buffer(data).Bytes())) {
		return
	}
	if !mycheck.Error2(f.WriteString("\n")) {
		return
	}
	return mycheck.Error(f.Close())
}
func (o *object) WriteGoCode(name string, data any) (ok bool) {
	b, err := format.Source(o.buffer(data).Bytes())
	if !mycheck.Error(err) {
		return
	}
	return o.WriteTruncate(name, b)
}
func (o *object) WriteBinary(name string, data any) (ok bool) { return o.WriteTruncate(name, data) }
func (o *object) ToLines(data any) (lines []string, ok bool) {
	newReader := bufio.NewReader(o.buffer(data))
	for {
		line, _, err := newReader.ReadLine()
		switch err {
		case io.EOF:
			return lines, true
		default:
			if !mycheck.Error(err) {
				return
			}
		}
		lines = append(lines, string(line))
	}
}
func (o *object) WriteJson(name string, Obj any) (ok bool) {
	var oo any
	switch reflect.TypeOf(Obj).Kind() {
	case reflect.Struct:
		oo = &Obj
	case reflect.Ptr:
		oo = Obj
	}
	data, err := json.MarshalIndent(oo, " ", " ")
	if !mycheck.Error(err) {
		return
	}
	return o.WriteTruncate(name, data)
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
	if !mycheck.Error(err) {
		return
	}
	return o.WriteTruncate(name, data)
}
func (o *object) GoCode() string { return o.goCode }
