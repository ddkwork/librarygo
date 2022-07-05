package client

import (
	"encoding/json"
	"github.com/ddkwork/librarygo/src/check"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/net/clientAndServer/internal/packet"
	"github.com/ddkwork/librarygo/src/stream"
	"net"
)

var receive = make(chan *stream.Buffer) //todo sync

type (
	Interface interface { //stickyBag() //by short connection way
		Connect(address string) bool
		SendJson(objectPtr interface{}) (ok bool)                      //call send
		SendJsonWithHead(head string, objectPtr interface{}) (ok bool) //call send
		SendWithHead(head, body *stream.Buffer) (ok bool)              //call send
		Send(s *stream.Buffer) (ok bool)
		Receive() *stream.Buffer
		//MarshalIndent(objectPtr interface{}) *stream.Buffer
	}
	object struct {
		check   check.Interface
		conn    net.Conn
		data    packet.Interface
		err     error
		address string
		ok      bool
	}
)

func (o *object) MarshalIndent(objectPtr interface{}) *stream.Buffer {
	send, err := json.MarshalIndent(objectPtr, " ", " ")
	o.ok = o.check.Error(err)
	return stream.NewBytes(send)
}

func (o *object) SendJsonWithHead(head string, objectPtr interface{}) (ok bool) {
	s := stream.NewString(head)
	marshalIndent := o.MarshalIndent(objectPtr)
	if !o.ok {
		return
	}
	return o.SendWithHead(s, marshalIndent)
}

func (o *object) SendJson(objectPtr interface{}) (ok bool) {
	marshalIndent := o.MarshalIndent(objectPtr)
	if !o.ok {
		return
	}
	return o.Send(marshalIndent)
}

func (o *object) SendWithHead(head, body *stream.Buffer) (ok bool) {
	buffer := stream.NewNil()
	buffer.Append(head, body)
	return o.Send(buffer)
}

func (o *object) stickyBag() bool {
	if !check.Error(o.conn.Close()) {
		return false
	}
	return o.Connect(o.address)
}

func (o *object) Connect(address string) bool {
	o.address = address
	o.conn, o.err = net.Dial("tcp", address)
	if !o.check.Error(o.err) {
		return false
	}
	mylog.Info("Client send to", address)
	return true
}

func (o *object) Receive() *stream.Buffer {
	return <-receive
}

func (o *object) Send(s *stream.Buffer) (ok bool) {
	if !o.check.Error2(o.conn.Write(s.Bytes())) {
		return false
	}
	mylog.Json("Client send stream", s.String())
	go func() {
		for {
			//o.data.Reset()
			n, err := o.conn.Read(o.data.Bytes())
			if err != nil || n == 0 {
				continue
			}
			data := make([]byte, n)
			copy(data, o.data.Bytes()[:n])
			s := stream.NewBytes(data)
			receive <- s //for sync every message
			if !o.stickyBag() {
				//todo handle sticky bag
			}
		}
	}()
	return true
}

func New() Interface {
	return &object{
		check: check.Default,
		conn:  nil,
		data:  packet.New(),
		err:   nil,
	}
}

var Default = New()
