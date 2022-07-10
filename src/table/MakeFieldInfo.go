package table

import (
	"bytes"
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"go/format"
	"strconv"
	"strings"
)

type FieldInfo struct {
	//Index     string
	Number    string
	Deep      string
	Name      string
	Kind      string
	Size      string
	ValueFmt  string //Hex|Decimal Bool String (HexDump Map)
	Value     any
	Stream    string
	Json      string //no
	TableLine *FieldInfo
}

func makeField(fieldInfos []FieldInfo) (code string) {
	b := new(bytes.Buffer)
	fnNewLine := func() { b.WriteByte('\n') }
	b.WriteString(`
func main() {
	b := new(bytes.Buffer)
	Style = DoubleSingleHorizontalBoxStyle
	Writer = b
	Run(
`)
	b.WriteString(strconv.Quote(
		"index\t"+
			"number\t"+
			"deep\t"+
			"name\t"+
			"kind\t"+
			"size\t"+
			"value\t"+
			"stream\n",
	) + "+")
	fnNewLine()
	for i, line := range fieldInfos {
		//if len(line.ValueFmt) > 23 {
		//	line.ValueFmt = wrap.Wrap(line.ValueFmt, 23)
		//	println(line.ValueFmt)
		//}
		if len(line.ValueFmt) > 23 {
			builder := strings.Builder{}
			builder.WriteString(line.ValueFmt[:10])
			builder.WriteString(" ... ")
			builder.WriteString(line.ValueFmt[len(line.ValueFmt)-10:])
			line.ValueFmt = builder.String()
		}

		lineText :=
			fmt.Sprint(i+1) + "\t" +
				line.Number + "\t" +
				line.Deep + "\t" +
				line.Name + "\t" +
				line.Kind + "\t" +
				line.Size + "\t"
		fnMakeSpace := func(count int) string { return strings.Repeat("\t", count) }
		if strings.Contains(line.ValueFmt, "\n") {
			line.ValueFmt = "\n" + line.ValueFmt
			lines, ok := tool.File().ToLines(bytes.NewBufferString(line.ValueFmt))
			if !ok {
				return
			}
			for _, s := range lines {
				lineText += fnMakeSpace(6) + s + "\t\n"
			}
		} else {
			lineText += line.ValueFmt + "\t"
		}
		lineText += line.Stream + "\n"
		lineText = strconv.Quote(lineText)
		if i+1 == len(fieldInfos) {
			lineText += ","
		} else {
			lineText += "+"
		}
		b.WriteString(lineText)
		fnNewLine()
	}
	b.WriteString(`false,
        Centred,
        Centred,
        Centred,
		LeftJustified,
		LeftJustified,
		Centred,
		LeftJustified,
		LeftJustified,
`) //少一个都会导致hexdump对不齐，不知道是什么原因
	b.WriteString(`)
}`)
	fnNewLine()
	source, err := format.Source(b.Bytes())
	if !mycheck.Error(err) {
		return
	}
	return string(source)
}
