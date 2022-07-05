package goBinary

import (
	"bytes"
	"encoding/gob"
	"github.com/ddkwork/librarygo/src/check"
)

type (
	Interface interface {
		Encode(obj interface{}) bool
		Decode(buf []byte, obj interface{}) bool
	}
	object struct {
		bytes.Buffer
		err error
	}
)

var c = check.Default

func New() Interface {
	return &object{}
}

func (o *object) Encode(obj interface{}) bool {
	enc := gob.NewEncoder(&o.Buffer)
	return c.Error(enc.Encode(obj))
}

func (o *object) Decode(buf []byte, obj interface{}) bool {
	if !c.Error2(o.Write(buf)) {
		return false
	}
	dec := gob.NewDecoder(&o.Buffer)
	return c.Error(dec.Decode(obj))
}
