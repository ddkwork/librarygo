package cpp2go

import (
	"context"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/ast"
	"github.com/goplus/c2go/clang/parser"
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

	p := "./Headers/Events.h"
	//D:\codespace\workspace\src\cppkit\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\forwarding.cpp
	c := "clang -Xclang -ast-dump=json -fsyntax-only "
	//c := "clang -cxx-isystem ./Headers -Xclang -ast-dump=json -fsyntax-only "
	session := NewSession()
	session.ShowLog = true
	session.SetDir(".")
	_, err2 := session.Run(context.Background(), c+p, true)
	if !mycheck.Error(err2) {
		return
	}
	//mylog.Json("ast", string(b))
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
