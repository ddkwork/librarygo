package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/ast"
	"github.com/goplus/c2go/clang/parser"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	// #include "BasicTypes.h"
	//typedef unsigned short wchar_t;
	//typedef void *PVOID;
	//typedef void *PVOID64;

	//D:\codespace\workspace\src\cppkit\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\hprdbgctrl.cpp
	p := "./Headers/Events.h"
	//p := "./Headers/Constants.h"
	//p := "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\app\\hprdbgctrl.cpp"
	//D:\codespace\workspace\src\cppkit\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\forwarding.cpp
	//c := "clang -Xclang -ast-dump=json -fsyntax-only "
	//"D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\
	//c := `clang -Xclang -dD -E -ast-dump=json -fsyntax-only `
	c := `gcc -posix -E -dM - < `
	abs, err3 := filepath.Abs(p)
	if !mycheck.Error(err3) {
		return
	}
	abs = strconv.Quote(abs)
	//b, err2 := session.Run(c + abs)
	c = c
	b, err2 := cmd.Run("C:\\Windows\\System32\\PING.EXE www.baidu.com -t ")
	if !mycheck.Error(err2) {
		return
	}
	mylog.Json("ast", b)
	select {}
	return
	node, warning, err := parser.ParseFile(p, 0)
	if !mycheck.Error(err) {
		return
	}
	mylog.Info("", warning)
	//mylog.Struct(node)
	for _, n := range node.Inner {
		//if n.Type != nil {
		//mylog.Struct(*n.Type)
		//}
		//if n.Field != nil {
		//	dumpNode("Field", n.Field)
		//}
		//if n.Decl != nil {
		//	dumpNode("Decl", n.Decl)
		//}
		//if n.OwnedTagDecl != nil {
		//	dumpNode("OwnedTagDecl", n.OwnedTagDecl)
		//}
		//mylog.Info("Kind", n.Kind)
		switch n.Kind {
		case ast.RecordDecl:
			mylog.Info("type struct{")
		case ast.TypedefType:
			mylog.Info("Type rename")
		case ast.FieldDecl:
			mylog.Info("FieldDecl Type, in field") //todo field end is FieldDecl
		case ast.EnumDecl:
			mylog.Info("const(")
		case ast.EnumConstantDecl:
			mylog.Info("EnumConstantDecl field")

		}
		dumpNode("Inner", *n)
		if n.Inner != nil {
			for _, n2 := range n.Inner {
				if n2 != nil {
					dumpNode("Inner", *n2)
				}
			}
		}

		//if n.ArrayFiller != nil {
		//	for _, n2 := range n.ArrayFiller {
		//		if n2 != nil {
		//			dumpNode("ArrayFiller", n2)
		//		}
		//	}
		//}

		//if n.Kind == ast.EnumDecl {
		//dumpNode("Kind == ast.EnumDecl", n)
		//}
	}
}

func dumpNode(title string, n ast.Node) {
	if n.Name == "" {
		return
	}
	//mylog.Info("----------", "--------------------")
	mylog.Warning(title, "")
	//mylog.Info("Name", n.Name)
	//mylog.Info("Kind Inner", n.Kind)
	mylog.Info("Kind", n.Kind)

	//mylog.Info("Size", n.Size)
	//mylog.Info("IsBitfield", n.IsBitfield)
	s := n.Name
	s += "\t\t"
	if n.Type != nil {
		//mylog.Info("QualType", n.Type.QualType)
		s += s + n.Type.QualType
		//s += "\t\t"
	}
	if n.Loc != nil {
		//mylog.Info("Line", n.Loc.Line)
		s += " //" + fmt.Sprint(n.Loc.Line)
	}
	mylog.Success("code", s)
	//mylog.Info("----------", "--------------------\n")
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
