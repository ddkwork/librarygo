package goBinary

import (
	"bytes"
	"encoding/gob"
	"github.com/ddkwork/librarygo/src/mycheck"
)

type (
	Interface interface {
		Encode(obj any) bool
		Decode(buf []byte, obj any) bool
	}
	object struct {
		bytes.Buffer
		err error
	}
)

var c = mycheck.Default

func New() Interface {
	return &object{}
}

func (o *object) Encode(obj any) bool {
	enc := gob.NewEncoder(&o.Buffer)
	return c.Error(enc.Encode(obj))
}

func (o *object) Decode(buf []byte, obj any) bool {
	if !c.Error2(o.Write(buf)) {
		return false
	}
	dec := gob.NewDecoder(&o.Buffer)
	return c.Error(dec.Decode(obj))
}
