package stream

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"math/big"
	"strings"
)

func (b2 *Buffer) NewLine()           { b2.WriteString("\n") }
func (b2 *Buffer) QuoteWith(s string) { b2.WriteString(s) }
func (b2 *Buffer) WriteBytesLn(p []byte) {
	(b2.Write(p))
	b2.NewLine()
}
func (b2 *Buffer) WriteStringLn(s string) {
	b2.WriteString(s)
	b2.NewLine()
}
func (b2 *Buffer) Quote()                 { b2.WriteByte('"') }
func (b2 *Buffer) ObjectBegin()           { b2.WriteByte('{') }
func (b2 *Buffer) ObjectEnd()             { b2.WriteByte('}') }
func (b2 *Buffer) SliceBegin()            { b2.WriteByte('[') }
func (b2 *Buffer) SliceEnd()              { b2.WriteByte(']') }
func (b2 *Buffer) Indent(deep int) string { return strings.Repeat(" ", deep) }
func (b2 *Buffer) HexString() string      { return hex.EncodeToString(b2.Bytes()) }
func (b2 *Buffer) HexStringUpper() string { return fmt.Sprintf("%#X", b2.Bytes())[2:] }
func (b2 *Buffer) SizeCheck() bool {
	switch b2.Len() {
	case 0:
		return mycheck.Error("buffer len == 0")
	default:
		if b2.Len()%8 != 0 {
			return mycheck.Error(" len%8 != 0")
		}
	}
	return true
}
func (b2 *Buffer) ErrorInfo() string { return b2.String() }
func (b2 *Buffer) Append(buffer ...*Buffer) {
	b2.Reset()
	for _, b2 := range buffer {
		b2.WriteBytesLn(b2.Bytes())
	}
}
func (b2 *Buffer) BigNumXorWithAlign(arg1, arg2 []byte, align int) (xorStream []byte) {
	xor := new(big.Int).Xor(new(big.Int).SetBytes(arg1), new(big.Int).SetBytes(arg2))
	alignBuf := make([]byte, align-len(xor.Bytes()))
	switch len(xor.Bytes()) {
	case 0:
		xorStream = alignBuf
	case align:
		xorStream = xor.Bytes()
	default:
		xorStream = b2.Merge(alignBuf, xor.Bytes()).Bytes()
	}
	return
}

func (b2 *Buffer) InsertString(splitSize int, separate string) (s string) {
	b := new(strings.Builder)
	for i, v := range b2.String() {
		b.WriteRune(v)
		if (i+1)%splitSize == 0 {
			b.WriteString(separate)
		}
	}
	s = b.String()
	s = s[:b.Len()-1]
	return
}
func (b2 *Buffer) SplitBytes(size int) (blocks [][]byte) {
	blocks = make([][]byte, 0)
	quantity := b2.Len() / size
	remainder := b2.Len() % size
	i := 0
	for i = 0; i < quantity; i++ {
		blocks = append(blocks, b2.Bytes()[i*size:(i+1)*size])
	}
	if remainder != 0 {
		blocks = append(blocks, b2.Bytes()[i*size:i*size+remainder])
	}
	return
}

func (b2 *Buffer) Merge(bytesSlice ...[]byte) *Buffer {
	b := bytes.NewBuffer(nil)
	b.Write(b2.Bytes())
	for i := 0; i < len(bytesSlice); i++ {
		if !mycheck.Error2(b.Write(bytesSlice[i])) {
			return nil
		}
	}
	return NewBuffer(b)
}
