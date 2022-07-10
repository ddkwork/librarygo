package httpClient

import (
	"crypto/tls"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type helper interface { //辅助接口
	SetForm(form url.Values) //单提交的两种方式：NewRequest 和 PostForm,选择NewRequest，因为PostForm没有设置请求头的接口
	SetRequestBody(requestBody []byte)
	Client() *http.Client
	CreatNewClient(client *http.Client)
	SetProxy(protocol, hostPort string) (ok bool)
	CheckProtocol(Protocol, hostPort string) (ok bool)
	Cookiejar() *cookiejar.Jar //自动管理原理猜测：返回的头自动添加到当前客户端的请求头
	SetStopCode(stopCode int)  //个别爬虫不是在返回200的时候读取数据的
	SetHead(head map[string]string)
	SetPath(path string)
	BaseURL() string
	SetRequestUrl(RequestUrl string)
	SetMethod(method string)
	SetMethodGet()
	SetMethodPost()
	hasCookieInJar(jar *cookiejar.Jar, cookieName, Host string) (ok bool) //另类网站的爬虫需要判断
}

//辅助:单一原则,get和set权衡一波
func (o *object) SetMethod(method string)            { o.method = method }
func (o *object) SetMethodGet()                      { o.method = http.MethodGet }
func (o *object) SetMethodPost()                     { o.method = http.MethodPost }
func (o *object) SetRequestUrl(RequestUrl string)    { o.requestUrl = RequestUrl }
func (o *object) SetPath(path string)                { o.path = path }
func (o *object) SetHead(head map[string]string)     { o.head = head }
func (o *object) SetStopCode(stopCode int)           { o.stopCode = stopCode }
func (o *object) BaseURL() string                    { return o.requestUrl }
func (o *object) Cookiejar() *cookiejar.Jar          { return o.cookiejar }
func (o *object) SetForm(form url.Values)            { o.form = form }
func (o *object) SetRequestBody(requestBody []byte)  { o.requestBody = requestBody }
func (o *object) Client() *http.Client               { return o.client }
func (o *object) CreatNewClient(client *http.Client) { o.client = client }
func (o *object) hasCookieInJar(jar *cookiejar.Jar, cookieName, Host string) (ok bool) {
	URL, err := url.Parse(Host)
	if !mycheck.Error(err) {
		return
	}
	for _, v := range jar.Cookies(URL) {
		if v.Name == cookieName {
			mylog.Success(" find cookie by name", v)
			return true
		}
	}
	return
}
func (o *object) SetProxy(protocol, hostPort string) (ok bool) { //todo see x/proxy pkg
	var ObjTransport struct {
		Transport    *http.Transport
		dialFunc     func(network, addr string) (net.Conn, error)
		proxyURLFunc func(*http.Request) (*url.URL, error)
	}
	if !o.CheckProtocol(protocol, hostPort) {
		return
	}
	switch protocol {
	//理论参数:
	//
	//
	//
	//
	case ProtoName.Socks4(), ProtoName.Socks5():
		ObjTransport.dialFunc = SDial(protocol + "://" + hostPort + "?timeout=20s") //?
	case ProtoName.Http(), ProtoName.Https():
		URL, err := url.Parse(ProtoName.Http() + "://" + hostPort)
		if !mycheck.Error(err) {
			return
		}
		ObjTransport.proxyURLFunc = http.ProxyURL(URL)
		ObjTransport.dialFunc = (&net.Dialer{
			Timeout:       6 * time.Second,                  //代理ip建立连接的话6秒就行了，因为下面的MaxIdleConns 20个链接指定可以换20个ip*6=120秒了，大于客户端超时60秒了
			Deadline:      time.Now().Add(20 * time.Second), //当前代理ip把数据取回来指定在20秒内，取不回来就换下一个ip
			LocalAddr:     nil,
			DualStack:     false,
			FallbackDelay: 0, //取到数据后延时一下？如果为零，则使用 300 毫秒的默认延迟，不延时的化服务器压测会使得很多代理误判师失效
			KeepAlive:     0, // 15 秒）发送保持活动探测。不支持 keep-alives 的网络协议或操作系统会忽略此字段。如果为负
			Resolver:      nil,
			Cancel:        nil,
			Control:       nil,
		}).Dial
	}
	o.client.Transport = &http.Transport{
		Proxy:                  ObjTransport.proxyURLFunc,
		DialContext:            nil,
		Dial:                   ObjTransport.dialFunc,
		DialTLSContext:         nil, //如果抓https包调这个封装的话旧的那个了
		DialTLS:                nil,
		TLSClientConfig:        &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      true, //禁用端口转发连接池,这个设置会全局保持超时控制的合理性
		DisableCompression:     false,
		MaxIdleConns:           10, //每个代理连接的超时 * 个数应该<客户端超市，否则该次网络请求就超时了，考虑到当前代理ip无效，可以更换10个ip * 30=300秒 >客户端请求超时60秒
		MaxIdleConnsPerHost:    10,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        0, //空闲连接超时
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0, //弄好代理参数立马发出去，还等个毛啊
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}
	return true
}
