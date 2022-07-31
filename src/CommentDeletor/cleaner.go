package CommentDeletor

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool/file"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (o *object) Exts() []string {
	return []string{
		".c",
		".cc",
		".cpp",
		".cppm",
		".h",
		".hh",
		".hpp",
		".ixx",
		".cs",
		".go",
	}
}
func (o *object) hasExt() (ok bool) {
	for _, s := range o.Exts() {
		if s == filepath.Ext(o.path) {
			fmt.Println(o.path)
			return true
		}
	}
	return
}
func (o *object) FileToLines(path string) (ok bool) {
	o.path = path
	b, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	lines, ok := file.New().ToLines(b)
	if !ok {
		return
	}
	o.lines = lines
	return true
}
func (o *object) debug() {
	for _, line := range o.skipLines {
		fmt.Printf("file-->%s  line-->%04d code-->%s", o.path, line.index, strconv.Quote(line.code))
		fmt.Println()
	}
	//for _, line := range o.lines {
	//	fmt.Println(line)
	//}
}
func (o *object) ResetLine() {
	for _, line := range o.skipLines {
		o.lines[line.index] = ""
	}
}
func (o *object) Size() int { return len(o.lines) }
func (o *object) startEnd(line string) (noSpace, start, end string) {
	noSpace = strings.TrimSpace(line)
	if len(noSpace) > 1 {
		start = noSpace[:2]
		end = noSpace[len(noSpace)-2:]
	}
	return
}
