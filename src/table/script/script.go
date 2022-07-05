package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b := new(bytes.Buffer)
	Style = DoubleSingleHorizontalBoxStyle
	Writer = b
	Run(
		"index\tnumber\tdeep\tname\tkind\tsize\tvalue\tstream\n"+
			"1\t1\t1\tGroup1\tObjectStart\t\t{\t\n"+
			"2\t1\t1\tBinary1\tstring\t24\tgame/system/session/info\t\n",
		false,
		Centred,
		Centred,
		Centred,
		LeftJustified,
		LeftJustified,
		Centred,
		LeftJustified,
		LeftJustified,
	)
}

func Run(tabulated string, isDoc bool, cellPrinters ...func(string, int)) string {
	Print(tabulated, cellPrinters...)
	b := Writer.(*bytes.Buffer)
	if !isDoc {
		println(b.String())
		return b.String()
	}
	lines, ok := ToLines(b)
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

func ToLines(data *bytes.Buffer) (lines []string, ok bool) {
	newReader := bufio.NewReader(data)
	for {
		line, _, err := newReader.ReadLine()
		switch err {
		case io.EOF:
			return lines, true
		default:
			if err != nil {
				return
			}
		}
		lines = append(lines, string(line))
	}
}

// built-in Styles
const (
	ASCIIStyle                            string = "+-++| ||"
	ASCIIBoxStyle                         string = "+-+++-++| ||+-++"
	MarkdownStyle                         string = "|-||| ||"
	BoxStyle                              string = "┌─┬┐├─┼┤│ ││└─┴┘"
	DoubleBoxStyle                        string = "╔═╦╗╠═╬╣║ ║║╚═╩╝"
	ThickHeaderDivideBoxStyle             string = "┌─┬┐┝━┿┥│ ││└─┴┘"
	DoubleBorderBoxStyle                  string = "╔═╤╗╟─┼╢║ │║╚═╧╝"
	DoubleVerticalBoxStyle                string = "╓─╥╖╟─╫╢║ ║║╙─╨╜"
	DoubleHorizontalBoxStyle              string = "╒═╤╕╞═╪╡│ ││╘═╧╛"
	DoubleSingleHorizontalBoxStyle        string = "╔═╤╗╠═╪╣║ │║╚═╧╝"
	DoubleTopBottomBoxStyle               string = "╒═╤╕├─┼┤│ ││╘═╧╛"
	DoubleSidesBoxStyle                   string = "╓─┬╖╟─┼╢║ │║╙─┴╜"
	DoubleTopBoxStyle                     string = "╒═╤╕├─┼┤│ ││└─┴┘"
	DoubleDivideBoxStyle                  string = "┌─┬┐╞═╪╡│ ││└─┴┘"
	DoubleBottomBoxStyle                  string = "┌─┬┐├─┼┤│ ││╘═╧╛"
	DoubleRightBoxStyle                   string = "┌─┬╖├─┼┤│ ││└─┴╜"
	DoubleLeftBoxStyle                    string = "╓─┬┐╟─┼┤║ ││╙─┴┘"
	DoubleInsideBoxStyle                  string = "┌─╥┐╞═╬╡│ ║│└─╨┘"
	DoubleInsideVerticalBoxStyle          string = "┌─╥┐├─╫┤│ ║│└─╨┘"
	DoubleInsideHorizontalBoxStyle        string = "┌─┬┐╞═╪╡│ ││└─┴┘"
	RoundedBoxStyle                       string = "╭─┬╮├─┼┤│ ││╰─┴╯"
	RoundedDoubleInsideBoxStyle           string = "╭─╥╮╞═╬╡│ ║│╰─╨╯"
	RoundedDoubleInsideHorizontalBoxStyle string = "╭─┬╮╞═╪╡│ ││╰─┴╯"
	RoundedDoubleInsideVerticalBoxStyle   string = "╭─╥╮├─╫┤│ ║│╰─╨╯"
)

// global options
var (
	Writer              io.Writer
	HeaderRows          = 1
	Style               = MarkdownStyle
	ColumnMapper        func(int) int // rearrange columns
	SortColumn          int
	NumericNotAlphaSort bool
	DefaultCellPrinter  = Centred
	DividerEvery        int
	FormfeedWithDivider bool
)

type codePoint []byte

// write a code point a number of times
func (c codePoint) repeat(w int) {
	for i := 0; i < w; i++ {
		Writer.Write(c)
	}
}

var cellPrinterPadding codePoint

type rowStyling struct {
	left, padding, divider, right codePoint
}

// set global var 'Writer' then call Print.
func Fprint(w io.Writer, tabulated string, cellPrinters ...func(string, int)) {
	Writer = w
	Print(tabulated, cellPrinters...)
}

// set global var 'Style' then call Print.
func Printf(s string, tabulated string, cellPrinters ...func(string, int)) {
	Style = s
	Print(tabulated, cellPrinters...)
}

// set the global var's 'Writer' and 'Style' then call Print.
func Fprintf(w io.Writer, s string, tabulated string, cellPrinters ...func(string, int)) {
	Writer = w
	Style = s
	Print(tabulated, cellPrinters...)
}

// Write 'tabulated' string as text table, rows coming from lines, columns separated by the tab character.
// Mono-spaced font required for alignment.
// cellPrinters - applied to columns:
// * missing - use default
// * len=1 - use for all cells
// * len=n - use n'th for n'th column, use default if column count>n
// Not thread safe, uses globals for options (see variables), however can be used multiple, fixed count, times by using multiple imports and different aliases.
// Bytes supporting.
//将“制表”字符串写为文本表，行来自行，列由制表符分隔。对齐所需的等距字体。 cellPrinters - 应用于列：缺失 - 使用默认值 len=1 - 用于所有单元格 len=n - 对第 n 列使用 n'th，如果列数>n 则使用默认值 不是线程安全的，使用全局变量作为选项（参见变量)，但是可以通过使用多个导入和不同的别名来使用多个固定计数的时间。 Bytes 支持。
func Print(tabulated string, cellPrinters ...func(string, int)) {
	// find max rows/widths, record cell strings 找到最大行宽，记录单元格字符串
	var columnMaxWidths []int
	var cells [][]string
	lineScanner := bufio.NewScanner(strings.NewReader(tabulated))
	for lineScanner.Scan() {
		rowCells := strings.Split(lineScanner.Text(), "\t")
		if needed := len(rowCells) - len(columnMaxWidths); needed > 0 {
			columnMaxWidths = append(columnMaxWidths, make([]int, needed)...)
		}
		for ci := range rowCells {
			if len(rowCells[ci]) > columnMaxWidths[ci] {
				columnMaxWidths[ci] = len(rowCells[ci])
			}
		}
		// order by function 按功能排序
		cells = append(cells, rowCells)
	}

	// order sortColumn 排序列
	if SortColumn > 0 {
		if HeaderRows < len(cells) {
			if HeaderRows < 0 {
				if NumericNotAlphaSort {
					sort.Sort(byColumnNumeric{byColumn{cells}})
				} else {
					sort.Sort(byColumnAlpha{byColumn{cells}})
				}
			} else {
				if NumericNotAlphaSort {
					sort.Sort(byColumnNumeric{byColumn{cells[HeaderRows:]}})
				} else {
					sort.Sort(byColumnAlpha{byColumn{cells[HeaderRows:]}})
				}
			}
		}
	}

	// the cellPrinter needed for a column 列所需的 cellPrinter
	cellPrinter := func(c int) func(string, int) {
		if len(cellPrinters) == 1 {
			return cellPrinters[0]
		}
		if c < len(cellPrinters) {
			return cellPrinters[c]
		}
		return DefaultCellPrinter
	}

	// use a scanner to split Style string into individual UTF8 code points 使用扫描仪将 Style 字符串拆分为单独的 UTF8 代码点
	runeScanner := bufio.NewScanner(strings.NewReader(Style))
	runeScanner.Split(bufio.ScanRunes)

	// scan a row style, 4 code points. 扫描一行样式，4个码点。
	scanRowStyling := func() *rowStyling {
		rf := new(rowStyling)
		runeScanner.Scan()
		rf.left = codePoint(runeScanner.Bytes())
		runeScanner.Scan()
		rf.padding = codePoint(runeScanner.Bytes())
		runeScanner.Scan()
		rf.divider = codePoint(runeScanner.Bytes())
		if !runeScanner.Scan() {
			return nil
		}
		rf.right = codePoint(runeScanner.Bytes())
		return rf
	}

	// write a content-less row using a row style, do nothing if nil. 使用行样式编写无内容行，如果为零则不执行任何操作。
	// used for top/bottom border and divider rows 用于上下边框和分隔行
	writeRow := func(rf *rowStyling) {
		if rf == nil {
			return
		}
		Writer.Write(rf.left)
		cellPrinterPadding = rf.padding
		if ColumnMapper == nil {
			for column, width := range columnMaxWidths {
				cellPrinter(column)("", width)
				if column < len(columnMaxWidths)-1 {
					Writer.Write(rf.divider)
				}
			}
		} else {
			for column := range columnMaxWidths {
				c := ColumnMapper(column)
				cellPrinter(column)("", columnMaxWidths[c])
				if column < len(columnMaxWidths)-1 {
					Writer.Write(rf.divider)
				}
			}
		}
		Writer.Write(rf.right)
		fmt.Fprintln(Writer)
	}

	// scan and store row Stylings from Style string, use helpful assumptions when not all blocks present. 从样式字符串中扫描并存储行样式，当并非所有块都存在时使用有用的假设。
	var dividerRowStyling, cellRowStyling, topRowStyling *rowStyling
	firstRowStyling := scanRowStyling()
	if firstRowStyling == nil {
		fmt.Fprintf(Writer, "Style %s needs to have at least 4 characters.", Style) //至少需要 4 个字符
		return
	}
	secondRowStyling := scanRowStyling()
	if secondRowStyling == nil {
		secondRowStyling = firstRowStyling
	}
	thirdRowStyling := scanRowStyling()
	if thirdRowStyling == nil {
		dividerRowStyling = firstRowStyling
		cellRowStyling = secondRowStyling
		topRowStyling = nil
	} else {
		dividerRowStyling = secondRowStyling
		cellRowStyling = thirdRowStyling
		topRowStyling = firstRowStyling
	}

	// write table 写表
	writeRow(topRowStyling)
	cellPrinterPadding = cellRowStyling.padding
	for row := range cells {
		if row-HeaderRows == 0 {
			writeRow(dividerRowStyling)
			cellPrinterPadding = cellRowStyling.padding
		}
		if DividerEvery > 0 && row-HeaderRows > DividerEvery && (row-HeaderRows)%DividerEvery == 0 {
			writeRow(dividerRowStyling)
			if FormfeedWithDivider {
				Writer.Write([]byte("\f"))
				writeRow(dividerRowStyling)
			}
			cellPrinterPadding = cellRowStyling.padding
		}
		Writer.Write(cellRowStyling.left)
		if ColumnMapper == nil {
			for column, cell := range cells[row] {
				cellPrinter(column)(cell, columnMaxWidths[column])
				if column < len(cells[row])-1 {
					Writer.Write(cellRowStyling.divider)
				}
			}
		} else {
			for column := range cells[row] {
				c := ColumnMapper(column)
				cellPrinter(column)(cells[row][c], columnMaxWidths[c])
				if column < len(columnMaxWidths)-1 {
					Writer.Write(cellRowStyling.divider)
				}
			}
		}
		Writer.Write(cellRowStyling.right)
		fmt.Fprintln(Writer)
	}
	// scan remaining row styling, if any, from Style for bottom border row. 从底部边框行的样式中扫描剩余的行样式（如果有）。
	writeRow(scanRowStyling())
}

// #cellPrinters 细胞打印机

// right justifier printer 右对齐打印机
func RightJustified(c string, w int) {
	cellPrinterPadding.repeat(w - len([]rune(c)))
	fmt.Fprint(Writer, c)
}

// left justifier printer 左对齐打印机
func LeftJustified(c string, w int) {
	fmt.Fprint(Writer, c)
	cellPrinterPadding.repeat(w - len([]rune(c)))
}

// centre printer 打印机中心
func Centred(c string, w int) {
	lc := len([]rune(c))
	offset := ((w - lc + 1) / 2)
	cellPrinterPadding.repeat(offset)
	fmt.Fprint(Writer, c)
	cellPrinterPadding.repeat(w - lc - offset)
}

// centre print if a boolean, right justify if a number, default otherwise. 如果是布尔值则居中打印，如果是数字则右对齐，否则默认。
func NumbersBoolJustified(c string, w int) {
	_, err := strconv.ParseBool(c)
	if err == nil {
		Centred(c, w)
		return
	}
	NumbersRightJustified(c, w)
}

// right justify if a number 如果一个数字右对齐
func NumbersRightJustified(c string, w int) {
	_, err := strconv.ParseInt(c, 10, 64)
	if err == nil {
		RightJustified(c, w)
		return
	}
	DefaultCellPrinter(c, w)
}

// modify a cellPrinter to have a minimum width 修改 cellPrinter 以具有最小宽度
func MinWidth(form func(string, int), min uint) func(string, int) {
	m := int(min)
	return func(s string, w int) {
		if w < m {
			form(s, m)
			return
		}
		form(s, w)
	}
}

// #sorters, implementing sort.Interface  分拣机，实现 sort.Interface

type byColumn struct {
	Rows [][]string //行
}

func (a byColumn) Len() int      { return len(a.Rows) }
func (a byColumn) Swap(i, j int) { a.Rows[i], a.Rows[j] = a.Rows[j], a.Rows[i] }

type byColumnAlpha struct {
	byColumn
}

func (a byColumnAlpha) Less(i, j int) bool { return a.Rows[i][SortColumn-1] < a.Rows[j][SortColumn-1] }

type byColumnNumeric struct {
	byColumn
}

func (a byColumnNumeric) Less(i, j int) bool {
	v1, err1 := strconv.ParseFloat(a.Rows[i][SortColumn-1], 64)
	v2, err2 := strconv.ParseFloat(a.Rows[j][SortColumn-1], 64)
	if err1 == nil && err2 == nil {
		return v1 < v2
	}
	return err1 == nil
}

// #mappers 映射器

// returns a column mapper func, that puts a particular column first, (columns start from 1), otherwise preserves order. 返回一个列映射器函数，它将特定列放在第一位（列从 1 开始），否则保留顺序。
func MoveToLeftEdge(column uint) func(int) int {
	c := int(column - 1)
	return func(n int) int {
		if n == 0 {
			return c
		}
		if n <= c {
			return n - 1
		}
		return n
	}
}
