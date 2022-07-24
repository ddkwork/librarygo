package gzip

import (
	"bytes"
	"compress/gzip"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"io/ioutil"
)

type (
	Interface interface {
		Decode(in []byte) *stream.Stream
	}
	object struct{ s *stream.Stream }
)

func New() Interface { return &object{s: stream.NewNil()} }

func (o *object) Decode(in []byte) *stream.Stream {
	c := mycheck.Default
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if !(c.Error(err)) {
		return nil
	}
	defer func() {
		if reader == nil {
			c.Error("gzipReader == nil")
			return
		}
		c.Error(reader.Close())
	}()
	all, err2 := ioutil.ReadAll(reader)
	if !mycheck.Error(err2) {
		return stream.NewErrorInfo(err2.Error())
	}
	return stream.NewBytes(all)
}
