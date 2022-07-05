package table_test

import (
	"bytes"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/ddkwork/librarygo/src/table"
	"strings"
	"testing"
)

//func TestObjInfo(t *testing.T) {
//	buf := []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
//	table.MakeObjectInfo([]table.FieldInfo{
//		{Name: "info", Kind: "int", Size: "8", Hex: "0X1122334455667788", Decimal: "1234605616436508552", HexDump: ""},
//		{Name: "infsdsdo", Kind: "int", Size: "8", Hex: "0X34455667788", Decimal: "3592025438088", HexDump: hex.Dump(buf)},
//		{Name: "infsdsdo", Kind: "int", Size: "8", Hex: "0X34455667788", Decimal: "3592025438088", HexDump: hex.Dump(buf)},
//		{Name: "info", Kind: "int", Size: "8", Hex: "0X1122334455667788", Decimal: "1234605616436508552", HexDump: ""},
//		{Name: "info", Kind: "int", Size: "8", Hex: "0X1122334455667788", Decimal: "1234605616436508552", HexDump: ""},
//		{Name: "info", Kind: "int", Size: "8", Hex: "0X1122334455667788", Decimal: "1234605616436508552", HexDump: ""},
//	})
//}

func TestGenInterface(t *testing.T) {
	b := new(bytes.Buffer)
	table.Style = table.DoubleSingleHorizontalBoxStyle
	table.Writer = b
	println(table.Run(
		"index\tname\tkind\tsize\thex\tdecimal\thexdump\n"+
			"1\tinfo\tint\t8\t0X1122334455667788\t1234605616436508552\t\n"+
			"2\tinfsdsdo\tint\t8\t0X34455667788\t3592025438088\t\n\t\t\t\t\t\t00000000  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000010  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000020  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t00000030  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t00000040  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t00000050  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000060  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000070  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t00000080  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t00000090  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t000000a0  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t000000b0  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t000000c0  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t000000d0  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t000000e0  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t000000f0  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000100  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000110  03 04 05                                          |...|\n"+
			"3\tinfsdsdo\tint\t8\t0X34455667788\t3592025438088\t\n\t\t\t\t\t\t00000000  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000010  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000020  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t00000030  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t00000040  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t00000050  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000060  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000070  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t00000080  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t00000090  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t000000a0  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t000000b0  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t000000c0  03 04 05 01 02 03 04 05  01 02 03 04 05 01 02 03  |................|\n\t\t\t\t\t\t000000d0  04 05 01 02 03 04 05 01  02 03 04 05 01 02 03 04  |................|\n\t\t\t\t\t\t000000e0  05 01 02 03 04 05 01 02  03 04 05 01 02 03 04 05  |................|\n\t\t\t\t\t\t000000f0  01 02 03 04 05 01 02 03  04 05 01 02 03 04 05 01  |................|\n\t\t\t\t\t\t00000100  02 03 04 05 01 02 03 04  05 01 02 03 04 05 01 02  |................|\n\t\t\t\t\t\t00000110  03 04 05                                          |...|\n"+
			"4\tinfo\tint\t8\t0X1122334455667788\t1234605616436508552\t\n"+
			"5\tinfo\tint\t8\t0X1122334455667788\t1234605616436508552\t\n"+
			"6\tinfo\tint\t8\t0X1122334455667788\t1234605616436508552\t\n",
		false,
		table.Centred,
		table.LeftJustified,
		table.LeftJustified,
		table.Centred,
		table.LeftJustified,
		table.LeftJustified,
		table.LeftJustified,
	))
}

func TestStackFmt(t *testing.T) {
	info := "reason\tapifield.FieldInfo{Number:0, Name:\"\", ValueFmt:interface {}(nil)} not handle kind:invalid\n" +
		"file\tD:/mod/myencoding/internal/apiprotocol/singular/sing ular.go\n" +
		"line\t89\n" +
		"name\tmyencoding/internal/apiprotocol/singular.(*superRecovery4).MarshalSingular\n" +
		"time\t2022-01-20 21:54:01 \n" +
		"goroutine\t2\n" +
		"\t               >>>>>>>>>>>>>>>>>>>>>>>>>>>> stack <<<<<<<<<<<<<<<<<<<<<<<<<<<\n"
	stack := `goroutine 18 [running]:
runtime/debug.Stack()
	C:/Program Files/Go/src/runtime/debug/stack.go:24 +0x65
libraryGo/src/check.(*object).setErrorInfo(0xc00001cc30, {0xa95220, 0xc000094ed0})
	D:/mod/libraryGo/src/check/ErrorCheck.go:258 +0xe6d
libraryGo/src/check.(*object).Error(0xc0000e82a0, {0xa95220, 0xc000094ed0})
	D:/mod/libraryGo/src/check/ErrorCheck.go:130 +0x2a
libraryGo/src/check.Error(...)
	D:/mod/libraryGo/src/check/ErrorCheck.go:12
myencoding/internal/apiprotocol/singular.(*superRecovery4).MarshalSingular(0xa94aa0, {0x0, {0xb18060, 0x0}, {0x0, 0x0}})
	D:/mod/myencoding/internal/apiprotocol/singular/singular.go:89 +0 x5ba
myencoding/internal/apiprotocol/Struct.(*object).marshalStruct(0xc0000b0f00, {0xabc1c0, 0xc0000a6a40})
	D:/mod/myencoding/internal/apiprotocol/Struct/Struct.go:82 +0x2ca 
myencoding/internal/apiprotocol/Struct.(*object).MarshalMessage(0xc000039f28, {0xabc1c0, 0xc0000a6a40})
`
	stack = strings.ReplaceAll(stack, "\n\t", " ---> ")
	lines, ok := tool.File().ToLines(stack)
	if !ok {
		return
	}
	for _, s := range lines {
		info += "\t" + s + "\n"
	}

	b := new(bytes.Buffer)
	fnNewLine := func() { b.WriteByte('\n') }
	table.Fprintf(b, table.BoxStyle, info,
		table.Centred,
		table.LeftJustified,
	)
	fnNewLine()
	println(b.String())
}
