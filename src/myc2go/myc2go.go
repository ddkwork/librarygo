package myc2go

import (
	"github.com/ddkwork/librarygo/src/myc2go/c2go"
	"github.com/ddkwork/librarygo/src/myc2go/c2go/cl"
	"github.com/ddkwork/librarygo/src/myc2go/c2go/clang/preprocessor"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/goplus/gox"
	"path/filepath"
)

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
