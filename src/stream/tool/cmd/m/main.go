package main

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
)

func main() {
	//b, err := cmd.Run("C:\\Windows\\System32\\PING.EXE www.baidu.com -t ")
	b, err := cmd.Run("ping www.baidu.com -t ")
	if !mycheck.Error(err) {
		return
	}
	mylog.Json("ast", b)
	select {}
}
