package stream

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"math/big"
	"sort"
	"strings"
)

func (b *Buffer) NewLine()           { b.WriteString("\n") }
func (b *Buffer) QuoteWith(s string) { b.WriteString(s) }
func (b *Buffer) WriteBytesLn(p []byte) {
	(b.Write(p))
	b.NewLine()
}
func (b *Buffer) WriteStringLn(s string) {
	b.WriteString(s)
	b.NewLine()
}
func (b *Buffer) Quote()                 { b.WriteByte('"') }
func (b *Buffer) ObjectBegin()           { b.WriteByte('{') }
func (b *Buffer) ObjectEnd()             { b.WriteByte('}') }
func (b *Buffer) SliceBegin()            { b.WriteByte('[') }
func (b *Buffer) SliceEnd()              { b.WriteByte(']') }
func (b *Buffer) Indent(deep int) string { return strings.Repeat(" ", deep) }
func (b *Buffer) HexString() string      { return hex.EncodeToString(b.Bytes()) }
func (b *Buffer) HexStringUpper() string { return fmt.Sprintf("%#X", b.Bytes())[2:] }
func (b *Buffer) SizeCheck() bool {
	switch b.Len() {
	case 0:
		return mycheck.Error("buffer len == 0")
	default:
		if b.Len()%8 != 0 {
			return mycheck.Error(" len%8 != 0")
		}
	}
	return true
}
func (b *Buffer) ErrorInfo() string { return b.String() }
func (b *Buffer) Append(buffer ...*Buffer) {
	b.Reset()
	for _, b2 := range buffer {
		b.WriteBytesLn(b2.Bytes())
	}
}
func (o *Buffer) BigNumXorWithAlign(arg1, arg2 []byte, align int) (xorStream []byte) {
	xor := new(big.Int).Xor(new(big.Int).SetBytes(arg1), new(big.Int).SetBytes(arg2))
	alignBuf := make([]byte, align-len(xor.Bytes()))
	switch len(xor.Bytes()) {
	case 0:
		xorStream = alignBuf
	case align:
		xorStream = xor.Bytes()
	default:
		xorStream = o.Merge(alignBuf, xor.Bytes()).Bytes()
	}
	return
}

func (o *Buffer) InsertString(splitSize int, separate string) (s string) {
	b := new(strings.Builder)
	for i, v := range o.String() {
		b.WriteRune(v)
		if (i+1)%splitSize == 0 {
			b.WriteString(separate)
		}
	}
	s = b.String()
	s = s[:b.Len()-1]
	return
}
func (o *Buffer) SplitBytes(splitSize int) (blocks [][]byte) {
	blocks = make([][]byte, 0)
	quantity := o.Len() / splitSize
	remainder := o.Len() % splitSize
	i := (0)
	for i = (0); i < quantity; i++ {
		blocks = append(blocks, o.Bytes()[i*splitSize:(i+1)*splitSize])
	}
	if remainder != 0 {
		blocks = append(blocks, o.Bytes()[i*splitSize:i*splitSize+remainder])
	}
	return
}

// JoinBytes ea safe need join playerid string type
func (o *Buffer) Merge(bytesSlice ...[]byte) *Buffer {
	b := bytes.NewBuffer(nil)
	for i := 0; i < len(bytesSlice); i++ {
		if !mycheck.Error2(b.Write(bytesSlice[i])) {
			return nil
		}
	}
	return NewBuffer(b)
}

// RemoveRepeatedElement 数组去重
func (o *Buffer) RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func (o *Buffer) RemoveRepeatedElementV2(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
