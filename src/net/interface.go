package net

import "github.com/ddkwork/librarygo/src/net/httpClient"

type (
	Interface interface {
		HttpClient() httpClient.Interface
		Xxobject() xxInterface
	}
	object struct {
		xxobject   xxInterface
		httpClient httpClient.Interface
	}
)

func (o *object) HttpClient() httpClient.Interface {
	return o.httpClient
}

func (o *object) Xxobject() xxInterface {
	return o.xxobject
}

var Default = New()

func New() Interface {
	return &object{
		xxobject:   xxNew(),
		httpClient: httpClient.New(),
	}
}
