package cpp2go

import (
	"testing"
)

func Test_cmd(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	sess.SetDir("/tmp")

	a, b := sess.Run("ls -l && du -h")

	//
	t.Log("ok...", a, b)

}

func Test_cmd1(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	a, b := sess.Run("./a.sh")

	//
	t.Log("ok...", a, b)

}

func Test_cmd2(t *testing.T) {

	sess := NewSession()
	sess.ShowLog = true

	a, b := sess.Run("echo 123444 > 1.log")

	//
	t.Log("ok...", a, b)

}
