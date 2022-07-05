package cmd

import (
	"bufio"
	"bytes"
	"github.com/ddkwork/librarygo/src/check"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

type charset string

const (
	UTF8    = charset("UTF-8")
	GB18030 = charset("GB18030")
)

type (
	//todo "github.com/PeterYangs/gcmd"  参考这个修复bug
	Interface interface {
		UTF82GBK(src string) *stream.Buffer                //UTF8转GBK
		GBK2UTF8(src []byte) string                        //GBK转UTF8
		CmdBuf2ChineseString(arg interface{}) (str string) //CmdBuf转ChineseString
		CmdRunWithCheck(arg string)                        //输出cmd执行结果
	}
	object struct {
	}
)

func New() Interface { return &object{} }

func (o *object) convertByte2String(byte []byte, charset charset) string {
	var str string
	switch charset {
	case GB18030:
		decodeBytes, err := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		if !check.Error(err) {
			return "err GB18030.NewDecoder()"
		}
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func (o *object) checkCmdResult(arg string, cmd *exec.Cmd) {
	outReader, err := cmd.StdoutPipe()
	if !check.Error(err) {
		return
	}
	errReader, err := cmd.StderrPipe()
	if !check.Error(err) {
		return
	}
	cmdReader := io.MultiReader(outReader, errReader)

	if !check.Error(cmd.Run()) { //run是同步，start是异步
		return
	}
	Stdin := bufio.NewScanner(cmdReader)
	for Stdin.Scan() {
		cmdRe := o.convertByte2String(Stdin.Bytes(), GB18030)
		cmdRe = strings.Replace(cmdRe, "\r\n", "", -1)
		mylog.Info("cmd命令 "+arg+" 返回:", cmdRe)
	}
	//if !check.Error(cmd.Wait()) {
	//	return
	//}
}

func (o *object) CmdRunWithCheck(arg string) {
	o.checkCmdResult(arg, exec.Command(arg))
	//o.checkCmdResult(arg, exec.Command("cmd", "/C", arg))
	//o.checkCmdResult(arg, exec.Command("cmd", "/c", "start "+arg))
}

func (o *object) CmdBuf2ChineseString(arg interface{}) (str string) {
	switch arg.(type) {
	case string:
		str = o.convertByte2String([]byte(arg.(string)), GB18030)
	case []byte:
		str = o.convertByte2String(arg.([]byte), GB18030)
	}
	return
}

// UTF82GBK : transform UTF8 rune into GBK byte array
func (o *object) UTF82GBK(src string) *stream.Buffer {
	GB18030 := simplifiedchinese.All[0]
	all, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
	if !check.Error(err) {
		return stream.NewErrorInfo(err.Error())
	}
	return stream.NewBytes(all)
}

// GBK2UTF8 : transform  GBK byte array into UTF8 string
func (o *object) GBK2UTF8(src []byte) string {
	GB18030 := simplifiedchinese.All[0]
	all, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	if !check.Error(err) {
		return ""
	}
	return string(all)
}
