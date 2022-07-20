package mylog

import (
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/librarygo/src/stream/indent"
	"github.com/ddkwork/librarygo/src/stream/tool/file"
	"github.com/ddkwork/librarygo/src/stream/tool/platform"
	"github.com/ddkwork/librarygo/src/stream/tool/time"
	"strings"
)

const (
	tagHex     = `[    hex]`
	tagHexDump = `[   dump]`
	tagJson    = `[   json]`
	tagStruct  = `[ struct]`
	tagInfo    = `[   info]`
	tagTrace   = `[  trace]`
	//tagError   = `[error]` //moved to mycheck pkg
	tagWarning = `[warning]`
	tagSuccess = `[success]`
)

func (o *object) HexDump(title string, msg any) {
	*o = object{
		tag:   tagHexDump,
		title: title,
		color: 93, //colorYellow
		msg:   hex.Dump(msg.([]byte)),
		body:  "",
		debug: o.debug,
	}
	o.do()
}

func (o *object) Hex(title string, msg any) {
	*o = object{
		tag:   tagHex,
		title: title,
		color: 36, //CyanString
		msg:   fmt.Sprintf("%#X", msg),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Info(title string, msg ...any) {
	*o = object{
		tag:   tagInfo,
		title: title,
		color: 96, //colorBlue
		msg:   fmt.Sprint(msg...),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Trace(title string, msg ...any) {
	*o = object{
		tag:   tagTrace,
		title: title,
		color: 94, //HiBlue
		msg:   fmt.Sprint(msg...),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Warning(title string, msg ...any) {
	*o = object{
		tag:   tagWarning,
		title: title,
		color: colorMagenta,
		msg:   fmt.Sprint(msg...),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Json(title string, msg ...any) {
	*o = object{
		tag:   tagJson,
		title: title,
		color: colorGreen,
		msg:   fmt.Sprint(msg...),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Success(title string, msg ...any) {
	*o = object{
		tag:   tagSuccess,
		title: title,
		color: colorGreen,
		msg:   fmt.Sprint(msg...),
		debug: o.debug,
	}
	o.do()
}

func (o *object) Struct(msg ...any) {
	*o = object{
		tag:   tagStruct,
		title: "",
		color: 92,                         //HiGreen
		msg:   fmt.Sprintf("%#v", msg...), //todo if see ptr need *
		debug: o.debug,
	}
	o.do()
}

func (o *object) do() bool { //no mycheck return
	if platform.New().IsAndroid() {
		return false
	}
	//2021-05-08 08:42:51 [STRC]                             | struct { a int; b string; c []uint8 }{a:89, b:"jhjsbdd", c:[]uint8{0x11, 0x22, 0x33, 0x44}}
	indentTitle := time.New().GetTimeNowString() + o.tag + indent.New().Left(o.title)
	switch o.tag {
	case tagJson, tagHexDump:
		indentTitle += "\n"
	}
	indexByte := strings.IndexByte(o.msg, '[')
	if indexByte == 0 { //很多小日志换行就那个了，那么不要换行了，直接删除他
		//o.msg = strings.Replace(o.msg, "[", "[\n", 1) //特殊处理，fmt.sprint不知道怎么把这个加上了
		//o.msg = strings.ReplaceAll(o.msg, "[", "") //pb klv 或者packed的时候怎么破?
		//o.msg = strings.ReplaceAll(o.msg, "]", "")
		b := []byte(o.msg)
		b = b[1 : len(b)-1]
		o.msg = string(b)
	}
	o.body = indentTitle + o.msg
	ColorBody := fmt.Sprintf(colorFormat, o.color, o.body)
	if o.debug {
		fmt.Println(ColorBody)
	}
	return file.New().WriteAppend("log.log", o.body)
}

const (
	colorFormat = "\x1b[1m\x1b[%dm%s\x1b[0m"
)

const (
	colorRed = uint8(iota + 31)
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
)

//{BlackString, "%s", nil, "\x1b[30m%s\x1b[0m"},
//{RedString, "%s", nil, "\x1b[31m%s\x1b[0m"},
//{GreenString, "%s", nil, "\x1b[32m%s\x1b[0m"},
//{YellowString, "%s", nil, "\x1b[33m%s\x1b[0m"},
//{BlueString, "%s", nil, "\x1b[34m%s\x1b[0m"},
//{MagentaString, "%s", nil, "\x1b[35m%s\x1b[0m"},
//{CyanString, "%s", nil, "\x1b[36m%s\x1b[0m"},
//{WhiteString, "%s", nil, "\x1b[37m%s\x1b[0m"},
//{HiBlackString, "%s", nil, "\x1b[90m%s\x1b[0m"},
//{HiRedString, "%s", nil, "\x1b[91m%s\x1b[0m"},
//{HiGreenString, "%s", nil, "\x1b[92m%s\x1b[0m"},
//{HiYellowString, "%s", nil, "\x1b[93m%s\x1b[0m"},
//{HiBlueString, "%s", nil, "\x1b[94m%s\x1b[0m"},
//{HiMagentaString, "%s", nil, "\x1b[95m%s\x1b[0m"},
//{HiCyanString, "%s", nil, "\x1b[96m%s\x1b[0m"},
//{HiWhiteString, "%s", nil, "\x1b[97m%s\x1b[0m"},
