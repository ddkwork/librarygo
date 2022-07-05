package table

import (
	"bytes"
	"fmt"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"strconv"
)

type (
	GenInterfaceDoc struct { //标题栏,完毕后全部注释+空两个空格才能被goland识别
		//index    int    //Centred
		Api      string //LeftJustified
		Function string //LeftJustified
		Note     string //LeftJustified
		Todo     string //LeftJustified 放到尾巴注释，因为无法对齐
		//align    func(string, int)
	}
)

func MakeDoc(linesObj []GenInterfaceDoc) {
	b := new(bytes.Buffer)
	fnNewLine := func() { b.WriteByte('\n') }
	b.WriteString(`
func TestGenInterface(t *testing.T) {
	b := new(bytes.Bytes)
	table.Style = table.DoubleSingleHorizontalBoxStyle
	table.Writer = b
	table.Run(
`)
	b.WriteString(strconv.Quote("index\t"+"api\t"+"function\t"+"note\t"+"todo\n") + "+")
	fnNewLine()
	for i, line := range linesObj {
		lineText := fmt.Sprint(i+1) + "\t" + line.Api + "\t" + line.Function + "\t" + line.Note + "\t" + line.Todo + "\n"
		lineText = strconv.Quote(lineText)
		if i+1 == len(linesObj) {
			lineText += ","
		} else {
			lineText += "+"
		}
		b.WriteString(lineText)
		fnNewLine()
	}
	b.WriteString(`true,
        table.Centred,
		table.LeftJustified,
		table.LeftJustified,
		table.LeftJustified,
		table.LeftJustified,
`)
	b.WriteString(`)
}`)
	fnNewLine()
	println(b.String())
}

func Run(tabulated string, isDoc bool, cellPrinters ...func(string, int)) string {
	Print(tabulated, cellPrinters...)
	b := Writer.(*bytes.Buffer)
	if !isDoc {
		println(b.String())
		return b.String()
	}
	lines, ok := tool.File().ToLines(b)
	if !ok {
		panic(ok)
	}
	println("//")
	for _, line := range lines {
		println("//  " + line)
	}
	println("//")
	return b.String()
}
