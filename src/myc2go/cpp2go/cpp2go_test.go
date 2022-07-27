package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/ast"
	"github.com/goplus/c2go/clang/parser"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	// #include "BasicTypes.h"
	//typedef unsigned short wchar_t;
	//typedef void *PVOID;
	//typedef void *PVOID64;
	//#pragma once

	//D:\codespace\workspace\src\cppkit\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\hprdbgctrl.cpp
	//p := "./Headers/Events.h"
	p := "./Headers/Constants.h"
	c := `clang -Xclang -dM -E -ast-dump=json -fsyntax-only `
	b, err2 := cmd.Run(c + p)
	if !mycheck.Error(err2) {
		return
	}
	node, warning, err := parser.ParseFile(p, 0)
	if !mycheck.Error(err) {
		return
	}
	mylog.Info("", warning)
	for _, n := range node.Inner {
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
	}
	//mylog.Json("ast", b)
	lines, ok := tool.File().ToLines(b)
	if !ok {
		return
	}
	for _, line := range lines {
		if strings.Contains(line, `#define`) {
			split := strings.Split(line, ` `)
			key := split[1]
			value := strings.Join(split[2:], ``)
			fmt.Printf("%-70s%s\n", key, value)
		}
	}
}

func dumpNode(title string, n ast.Node) {
	if n.Name == "" {
		return
	}
	mylog.Warning(title, "")
	mylog.Info("Kind", n.Kind)
	s := n.Name
	s += "\t\t"
	if n.Type != nil {
		s += s + n.Type.QualType
	}
	if n.Loc != nil {
		s += " //" + fmt.Sprint(n.Loc.Line)
	}
	mylog.Success("code", s)
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
