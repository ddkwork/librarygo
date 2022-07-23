package stream

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
)

type (
	Type interface {
		~string | ~[]byte | ~*bytes.Buffer //todo and test type rename
	}
	_Interface interface { //todo  合并tool包
		NewLine()
		Quote() //手动Quote字符串避免造成换行失效
		QuoteWith(s string)
		ObjectBegin()
		ObjectEnd()
		SliceBegin()
		SliceEnd()
		Indent(deep int) string
		WriteBytesLn(p []byte)
		WriteStringLn(s string)
		HexString() string
		HexStringUpper() string
		Append(buffer ...*Buffer)
		WriteXMakeBody(key string, values ...string)
		SizeCheck() bool
		ErrorInfo() string
		CutWithIndex(x, y int)
		BigNumXorWithAlign(arg1, arg2 []byte, align int) (xorStream []byte)
		Merge(Bytes ...[]byte) *Buffer
		InsertString(splitSize int, separate string) (s string)
		SplitBytes(size int) (blocks [][]byte)
		SplitString(size int) (blocks []string)
		RemoveHexDumpNewLine(dump string) (newDump string)
	}
	Buffer struct{ *bytes.Buffer }
)

func (b2 *Buffer) RemoveHexDumpNewLine(dump string) (newDump string) {
	//strings.TrimSuffix()
	panic("implement me")
}

func (b2 *Buffer) SplitString(size int) (blocks []string) {
	blocks = make([]string, 0)
	splitBytes := b2.SplitBytes(size)
	for _, splitByte := range splitBytes {
		blocks = append(blocks, string(splitByte))
	}
	return
}

func (b2 *Buffer) CutWithIndex(x, y int) {
	//TODO implement me
	panic("implement me")
}

var Default = New()

func New() *Buffer {
	return &Buffer{
		Buffer: &bytes.Buffer{},
	}
}
func NewBytes(b []byte) *Buffer           { return &Buffer{bytes.NewBuffer(b)} }
func NewBuffer(buf *bytes.Buffer) *Buffer { return &Buffer{buf} }
func NewString(s string) *Buffer          { return &Buffer{Buffer: bytes.NewBufferString(s)} }
func NewHexString(s string) (b *Buffer) {
	b = New()
	decodeString, err := hex.DecodeString(s)
	if !mycheck.Error(err) {
		b.WriteString(err.Error())
		return
	}
	b.Write(decodeString)
	return
}
func NewHexStringOrBytes(data any) (b *Buffer) {
	switch data.(type) {
	case string:
		return NewHexString(data.(string))
	case []byte:
		return NewBytes(data.([]byte))
	}
	return NewErrorInfo(fmt.Sprintf("%t\t", data))
}
func NewNil() *Buffer                 { return New() }
func NewErrorInfo(err string) *Buffer { return NewString(err) }
func newInterface() _Interface {
	return &Buffer{
		Buffer: &bytes.Buffer{},
	}
}
