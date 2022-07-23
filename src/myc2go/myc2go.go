package myc2go

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
)

/*
bug fix on windows
librarygo/src/c2go/clang/types/types.go
	//Long    = types.Typ[uintptr(types.Int32)+unsafe.Sizeof(0)>>3]  // int32/int64
	Long = types.Typ[uintptr(types.Int32)] // int32/int64
	//Ulong   = types.Typ[uintptr(types.Uint32)+unsafe.Sizeof(0)>>3] // uint32/uint64
	Ulong   = types.Typ[uintptr(types.Uint32)] // uint32/uint64
*/

func Run(path string) {
	abs, err := filepath.Abs(path)
	if !mycheck.Error(err) {
		return
	}
	cl.SetDebug(cl.DbgFlagAll)
	preprocessor.SetDebug(preprocessor.DbgFlagAll)
	gox.SetDebug(gox.DbgFlagInstruction) // | gox.DbgFlagMatch)
	c2go.Run("main", abs, c2go.FlagDumpJson, nil)
}
