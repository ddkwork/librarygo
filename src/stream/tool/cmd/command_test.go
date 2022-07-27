package cmd_test

import (
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"testing"
)

func Test_cmd(t *testing.T) {
	//cmd.SetDir("/tmp")
	a, b := cmd.Run("ls -l && du -h")
	t.Log("ok...", a, b)
}

func Test_cmd1(t *testing.T) {
	a, b := cmd.Run("./a.sh")
	t.Log("ok...", a, b)
}

func Test_cmd2(t *testing.T) {
	a, b := cmd.Run("echo 123444 > 1.log")
	t.Log("ok...", a, b)
}
