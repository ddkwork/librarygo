package structBytes

import (
	"bytes"
	"github.com/ddkwork/librarygo/src/check"
	"github.com/ddkwork/librarygo/src/stream/structBytes/goBinary"
)

var c = check.Default

type (
	Interface interface {
		StructBytes() []byte
		StructToBytes(obj interface{}) bool
		BytesToStruct(StructBytes []byte, obj interface{}) bool
		goBinary.Interface
	}
	object struct {
		*bytes.Buffer
		goBinary goBinary.Interface
	}
)

func New() Interface {
	return &object{
		Buffer:   nil,
		goBinary: goBinary.New(),
	}
}

func (o *object) StructBytes() []byte                { return o.Bytes() }
func (o *object) StructToBytes(obj interface{}) bool { return o.Write(obj) }
func (o *object) BytesToStruct(StructBytes []byte, obj interface{}) bool {
	return o.Read(StructBytes, obj)
}
func (o *object) Encode(obj interface{}) bool             { return o.goBinary.Encode(obj) }
func (o *object) Decode(buf []byte, obj interface{}) bool { return o.goBinary.Decode(buf, obj) }
