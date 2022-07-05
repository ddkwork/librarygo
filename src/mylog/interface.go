package mylog

type (
	Interface interface {
		HexDump(title string, msg interface{})    //hex buf todo support fn return []byte
		Hex(title string, msg interface{})        //hex value
		Info(title string, msg ...interface{})    //info
		Trace(title string, msg ...interface{})   //跟踪
		Warning(title string, msg ...interface{}) //警告
		Json(title string, msg ...interface{})    //pb json todo rename
		Success(title string, msg ...interface{}) //成功
		Struct(msg ...interface{})                //结构体 todo indent ,add title
		Body() string
		Msg() string
		SetDebug(debug bool)
	}
	object struct {
		tag   string
		title string
		color uint8
		msg   string
		body  string
		debug bool
	}
)

func (o *object) SetDebug(debug bool) { o.debug = debug }
func (o *object) Msg() string         { return o.msg }
func (o *object) Body() string        { return o.body }

func New() Interface {
	return &object{
		tag:   "",
		title: "",
		color: 0,
		msg:   "",
		body:  "",
		debug: true,
	}
}
func init() {
	Trace("--------- title ---------", "------------------ info ------------------")
}

var Default = New()

func HexDump(title string, msg interface{})    { Default.HexDump(title, msg) }
func Hex(title string, msg interface{})        { Default.Hex(title, msg) }
func Info(title string, msg ...interface{})    { Default.Info(title, msg) }
func Trace(title string, msg ...interface{})   { Default.Trace(title, msg) }
func Warning(title string, msg ...interface{}) { Default.Warning(title, msg) }
func Json(title string, msg ...interface{})    { Default.Json(title, msg) }
func Success(title string, msg ...interface{}) { Default.Success(title, msg) }
func Struct(msg ...interface{})                { Default.Struct(msg) }
func Body() string                             { return Default.Body() }
func Msg() string                              { return Default.Msg() }
func SetDebug(debug bool)                      { Default.SetDebug(debug) }
