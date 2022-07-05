package server

import (
	"github.com/ddkwork/librarygo/src/check"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/net/clientAndServer/internal/packet"
	"github.com/ddkwork/librarygo/src/stream"
	"net"
)

var receive = make(chan *stream.Buffer)

type (
	Interface interface {
		ListenAndServer(address string) bool
		Receive() *stream.Buffer
		Replay(data string) bool
	}
	object struct {
		check check.Interface
		data  packet.Interface
		l     net.Listener
		conn  net.Conn
		err   error
	}
)

func (o *object) Replay(data string) bool {
	return o.check.Error2(o.conn.Write([]byte(data)))
}

func (o *object) Receive() *stream.Buffer { return <-receive }

func (o *object) ListenAndServer(address string) bool {
	if !o.Listen(address) {
		return false
	}
	go o.Server()
	return true
}

func (o *object) Listen(address string) bool {
	o.l, o.err = net.Listen("tcp", address)
	if !o.check.Error(o.err) {
		return false
	}
	mylog.Info("Server Listen on", address)
	return true
}

func (o *object) Server() {
	//wg := sync.WaitGroup{}
	for {
		o.conn, o.err = o.l.Accept()
		if !o.check.Error(o.err) {
			continue
		}
		//o.data.Reset()
		//go func() {
		//    defer func() { check.Error(conn.Close()) }()
		n, err := o.conn.Read(o.data.Bytes())
		if err != nil || n == 0 {
			continue
		}
		data := make([]byte, n)
		copy(data, o.data.Bytes()[:n])
		s := stream.NewBytes(data)
		receive <- s
		//}()
	}
}

func New() Interface {
	return &object{
		check: check.Default,
		data:  packet.New(),
	}
}

var Default = New()
