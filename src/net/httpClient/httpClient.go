package httpClient

import (
	"bytes"
	"errors"
	"github.com/ddkwork/librarygo/src/mycheck"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	Interface interface { //业务接口
		NewRequest() (ok bool)
		ResponseBuf() []byte //todo no gzip,gzip,json interface
		Error() error        //错误
		helper
	}

	object struct {
		client      *http.Client      //get set
		proxy       *http.Transport   //set
		cookiejar   *cookiejar.Jar    //get
		form        url.Values        //set
		requestBody []byte            //set
		method      string            //set
		requestUrl  string            //set
		path        string            //set
		head        map[string]string //set
		stopCode    int               //set
		responseBuf []byte            //get
		error       error             //get
	}
)

//业务
func (o *object) NewRequest() (ok bool) { //todo 考虑一个携程带上一个新的实例化对象来执行这个函数
	var (
		request  = new(http.Request)
		response = new(http.Response)
	)

	//新建请求
	var body io.Reader
	if o.requestBody != nil {
		body = bytes.NewReader(o.requestBody)
	} else {
		body = strings.NewReader(o.form.Encode())
	}

	request, o.error = http.NewRequest(o.method, o.requestUrl+o.path, body)
	if !mycheck.Error(o.error) {
		return
	}

	request.Close = true //强制短链接
	//Request.Header.Add("Connection", "close")

	//设置请求头
	for k, v := range o.head {
		request.Header.Set(k, v)
	}

	//设置请方式
	//request.Method = o.method

	//发起请求
	response, o.error = o.client.Do(request)
	if !mycheck.Error(o.error) {
		return
	}

	//关闭返回的io接口
	defer func() {
		if response == nil {
			o.error = errors.New("response == nil")
			mycheck.Error(o.error)
			return
		}
		mycheck.Error(response.Body.Close())
	}()

	//读返回body
	switch response.StatusCode {
	case http.StatusOK, o.stopCode:
		o.responseBuf, o.error = ioutil.ReadAll(response.Body) //todo 外部判断gzip，是否可以提进来，就不用重复劳动了
		return mycheck.Error(o.error)
	default:
		o.error = errors.New(response.Status + " != StopCode " + strconv.Itoa(o.stopCode))
		return mycheck.Error(o.error)
	}
}
func (o *object) ResponseBuf() []byte { return o.responseBuf }
func (o *object) Error() error        { return o.error }

var Default = New()

func New() Interface {
	jar, err := cookiejar.New(nil)
	if !mycheck.Error(err) {
		return nil
	}
	return &object{
		client: &http.Client{
			Transport: nil,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Jar:     nil,
			Timeout: 60 * time.Second, //如果代理参数太大的话会超时的，这里应该增大，如何不用代理的，60秒都太大了
		},
		proxy:      nil,
		cookiejar:  jar,
		form:       nil,
		method:     "",
		requestUrl: "",
		path:       "",
		head:       nil,
		stopCode:   http.StatusOK,
	}
}
