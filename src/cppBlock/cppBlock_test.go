package cppBlock

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	println(strings.Index("} ZydisDecoderState;", "}"))
	println(strings.Contains("} ZydisDecoderState;", "}"))
	println(strings.Contains(`#elif defined(ZYAN_WINDOWS)`, `#define`))
}

func TestFindAll(t *testing.T) {
	p := "Decoder.back"
	lines, ok := tool.File().ReadToLines(p)
	if !ok {
		panic(111)
	}
	l := FindStruct(lines)
	for _, info := range l {
		mylog.Info(fmt.Sprint(info.Col), info.Line) //51 - 137
	}
	l = FindEnum(lines)
	for _, info := range l {
		mylog.Info(fmt.Sprint(info.Col), info.Line) //147 - 222
	}
}

func TestFindExtern(t *testing.T) {
	p := "dt-struct.cpp.back"
	lines, ok := tool.File().ReadToLines(p)
	if !ok {
		panic(111)
	}
	l := FindExtern(lines)
	for _, info := range l {
		mylog.Info(fmt.Sprint(info.Col), info.Line) //147 - 222
	}
}

func TestFindDefine(t *testing.T) {
	p := "Thread.h.back"
	lines, ok := tool.File().ReadToLines(p)
	if !ok {
		panic(111)
	}
	l := FindDefine(lines)
	for _, info := range l {
		mylog.Info(fmt.Sprint(info.Col), info.Line) //147 - 222
	}
}

func TestFindMethod(t *testing.T) {
	p := "common.cpp.back"
	lines, ok := tool.File().ReadToLines(p)
	if !ok {
		panic(111)
	}
	l := FindMethod(lines)
	for _, info := range l {
		mylog.Info(fmt.Sprint(info.Col), info.Line) //147 - 222
	}
}
