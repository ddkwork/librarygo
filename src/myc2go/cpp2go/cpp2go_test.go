package cpp2go

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/parser"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	p := "./Headers/Events.h"
	//"-Xclang", "-ast-dump=json", "-fsyntax-only",
	if !cmd.Run("clang -I ./Headers -Xclang -ast-dump=json -fsyntax-only " + p) {
		return
	}
	return

	node, warning, err := parser.ParseFile(p, 0)
	if !mycheck.Error(err) {
		return
	}
	mylog.Info("", warning)
	mylog.Struct(node)

	return
	Run(p, "main")
}

func Run(path, pkg string) {
	abs, err := filepath.Abs(path)
	if !mycheck.Error(err) {
		return
	}
	cl.SetDebug(cl.DbgFlagAll)
	preprocessor.SetDebug(preprocessor.DbgFlagAll)
	gox.SetDebug(gox.DbgFlagInstruction) // | gox.DbgFlagMatch)
	c2go.Run(pkg, abs, c2go.FlagDumpJson, nil)
}
