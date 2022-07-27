package cmd

import (
	"github.com/PeterYangs/gcmd"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
)

func Run(command string) (ok bool) {
	mylog.Info("command", command)
	buf, err := gcmd.Command(command).ConvertUtf8().Start()
	if !mycheck.Error(err) {
		return
	}
	mylog.Info("cmd out", buf)
	return true
}
func RunWithReturn(command string) (b []byte, err error) {
	mylog.Info("command", command)
	return gcmd.Command(command).ConvertUtf8().Start()
}
