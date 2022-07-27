package cpp2go

import (
	"context"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	// #include "BasicTypes.h"

	//typedef unsigned short wchar_t;
	//typedef void *PVOID;
	//typedef void *PVOID64;

	session := NewSession()

	p := "./Headers/Events.h"
	//c := "clang -Xclang -ast-dump=json -fsyntax-only "
	c := "clang -cxx-isystem ./Headers -Xclang -ast-dump=json -fsyntax-only "

	session.ShowLog = true
	session.SetDir(".")
	b, err2 := session.Run(context.Background(), c+p, true)
	if !mycheck.Error(err2) {
		return
	}
	mylog.Json("ast", string(b))

	//node, warning, err := parser.ParseFile(p, 0)
	//if !mycheck.Error(err) {
	//	return
	//}
	//mylog.Info("", warning)
	//mylog.Struct(node)
	//
	//return
	//Run(p, "main")
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
