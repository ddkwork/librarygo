package myc2go

import (
	"github.com/ddkwork/librarygo/src/caseconv"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	Interface interface {
		ConvertAll() (ok bool)
		C2go() (ok bool)
		Format() (ok bool)
		SetRoot(root []string)
		SetExt(ext []string)
		SetContains(contains []string)
		SetContainsNot(containsNot []string)
		SetSkip(skip string)
		SetReplaces(replaces map[string]string)
	}
	Setup struct {
		SetRoot        func() []string
		SetExt         func() []string
		SetContains    func() []string
		SetContainsNot func() []string
		SetSkip        func() string
		SetReplaces    func() map[string]string
	}
	object struct {
		contains    []string
		containsNot []string
		root        []string
		ext         []string
		apis        []string
		replaces    map[string]string
		files       map[string]string
		skip        string
		stream      *stream.Stream
	}
)

func NewSetup(setup Setup) Interface {
	return &object{
		contains:    setup.SetContains(),
		containsNot: setup.SetContainsNot(),
		root:        setup.SetRoot(),
		ext:         setup.SetExt(),
		apis:        nil,
		replaces:    nil,
		files:       nil, //goFilePath goFileBody
		skip:        setup.SetSkip(),
		stream:      stream.New(),
	}
}

func (o *object) HandleDefines(path string) (ok bool) {
	Constants, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	lines, ok := tool.File().ToLines(Constants)
	if !ok {
		return
	}
	ss := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, "SCRIPT_ENGINE_KERNEL_MODE") {
			ss = lines[i+1:]
			break
		}
	}
	consts := make([]string, 0)
	define := ""
	for i := 0; i < len(ss); i++ {
		switch {
		case strings.Contains(ss[i], "#define") && strings.Contains(ss[i], `\`):
			defineEnd := 0
			define = ss[i+defineEnd]
			for {
				defineEnd++
				define += ss[i+defineEnd]
				define = strings.ReplaceAll(define, `\`, ``)
				define = strings.Trim(define, " ")
				if define != "" {
					consts = append(consts, define)
				}
				if strings.Contains(ss[i+defineEnd], "#define") {
					i += defineEnd - 1
					break
				}
			}
		case strings.Contains(ss[i], "#define") && !strings.Contains(ss[i], `\`):
			define = ss[i]
			if define != "" {
				consts = append(consts, define)
			}
		}
	}
	o.stream.WriteStringLn("const(")
	for _, define := range consts {
		if strings.Contains(define, "sizeof(UINT32)") {
			define = strings.ReplaceAll(define, "sizeof(UINT32)", "4") //todo
		}
		if strings.Contains(define, "sizeof(DEBUGGER_REMOTE_PACKET)") {
			define = strings.ReplaceAll(define, "sizeof(DEBUGGER_REMOTE_PACKET)", "11") //todo
		}
		//println(define)
		all := strings.ReplaceAll(define, "#define", "")
		all = strings.TrimSpace(all)
		index := strings.Index(all, " ")
		key := all[:index]
		value := all[index:]
		key += "  =" + value
		o.stream.WriteStringLn(key)
	}
	o.stream.WriteStringLn(")")
	return true
}
func (o *object) HandleApis(path string) (ok bool) {
	Constants, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	apis, ok := tool.File().ToLines(Constants)
	if !ok {
		return
	}
	ss := make([]string, 0)
	for i, line := range apis {
		if strings.Contains(line, "funcs") { //todo
			ss = apis[i+1:]
			break
		}
	}
	lines := make([]string, 0)
	lines, b := tool.File().ToLines(o.stream.String())
	if !b {
		return
	}
	s := stream.New()
	for _, line := range lines {
		s.WriteStringLn(line)
		if strings.Contains(line, apiStart) {
			for _, api := range apis {
				s.WriteStringLn(api + `()(ok bool)`) //todo add // origname
			}
		}
	}
	o.stream.Reset()
	o.stream.Write(s.Bytes())
	return true
}
func (o *object) Embed(path, objectName string) (ok bool) {
	abs, err := filepath.Abs(path)
	if !mycheck.Error(err) {
		return
	}
	o.stream.WriteStringLn(`//go:embed ` + strconv.Quote(abs))
	o.stream.WriteStringLn(`var ` + objectName + `Buf string`)
	return true
}
func (o *object) PkgName(path string) (pkgName string) {
	pkgName = filepath.Base(filepath.Dir(path))
	if strings.Contains(pkgName, "-") {
		pkgName = strings.ReplaceAll(pkgName, "-", "")
	}
	if strings.Contains(pkgName, "~") {
		pkgName = strings.ReplaceAll(pkgName, "~", "unknown")
	}
	return
}

const (
	apiStart = `		//Fn() (ok bool)`
)

func (o *object) ObjectName(path string) (objectName, ext string) {
	ext = filepath.Ext(path)
	objectName = filepath.Base(path)
	objectName = objectName[:len(objectName)-len(ext)]
	if strings.Contains(objectName, `~`) {
		objectName = strings.ReplaceAll(objectName, `~`, `unknown`)
	}
	if strings.Contains(objectName, `-`) {
		objectName = strings.ReplaceAll(objectName, `-`, `_`)
	}
	if strings.Contains(objectName, `switch`) {
		objectName = strings.ReplaceAll(objectName, `switch`, `switchA`)
	}
	objectName = caseconv.ToCamel(objectName, false)
	objectName = strings.TrimRight(objectName, " ")
	return
}
func (o *object) ConvertAll() (ok bool) {
	for _, p := range o.root {
		if !mycheck.Error(filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
			for _, contain := range o.contains {
				if strings.Contains(path, contain) {
					for _, s2 := range o.containsNot {
						if !strings.Contains(path, s2) {
							objectName, ext := o.ObjectName(path)
							for _, e := range o.ext {
								if e == ext {
									o.stream.Reset()
									pkgName := o.PkgName(path)
									o.stream.WriteStringLn("package " + pkgName)
									o.stream.WriteStringLn(`import (_ "embed")`)
									if !o.Embed(path, objectName) {
										return err
									}
									o.stream.WriteStringLn(`type (`)
									o.stream.WriteStringLn(caseconv.ToCamelUpper(objectName, false) + ` interface {`)
									o.stream.WriteStringLn(apiStart)
									o.stream.WriteStringLn(`	}`)
									o.stream.WriteStringLn(caseconv.ToCamel(objectName, false) + `  struct{}`)
									o.stream.WriteStringLn(`)`)
									o.stream.WriteStringLn(`func New` + caseconv.ToCamel(objectName, false) + `() ` +
										caseconv.ToCamelUpper(objectName, false) + ` { return & ` + caseconv.ToCamel(objectName, false) + `{} }`)
									//buffer.WriteStringLn(``)
									//println(buffer.String())
									goFilePath := filepath.Join("go", filepath.Dir(path), objectName+".go")
									o.files[goFilePath] = o.stream.String()
									if !o.HandleApis(path) {
										return err
									}
									if !o.HandleDefines(path) {
										return err
									}
								}
							}
						}
					}
				}
			}
			return err
		})) {
			return
		}
	}
	return o.Format()
}

func (o *object) C2go() (ok bool) {
	for _, file := range o.files {
		Run(file)
	}
	return o.Format()
}

func (o *object) Format() (ok bool) {
	for path, body := range o.files {
		if !tool.File().WriteTruncate(path, body) {
			return
		}
	}
	for path := range o.files {
		readFile, err := os.ReadFile(path)
		if !mycheck.Error(err) {
			return
		}
		source, err := format.Source(readFile)
		if !mycheck.Error(err) {
			return
		}
		if !tool.File().WriteTruncate(path, source) {
			return
		}
	}
	return true
}
func (o *object) SetSkip(skip string)                    { o.skip = skip }
func (o *object) SetExt(ext []string)                    { o.ext = ext }
func (o *object) SetRoot(root []string)                  { o.root = root }
func (o *object) SetContainsNot(containsNot []string)    { o.containsNot = containsNot }
func (o *object) SetContains(contains []string)          { o.contains = contains }
func (o *object) SetReplaces(replaces map[string]string) { o.replaces = replaces }

func New() Interface {
	return &object{
		contains:    nil,
		containsNot: nil,
		root:        nil,
		ext:         nil,
		apis:        nil,
		replaces:    nil,
		files:       make(map[string]string),
		skip:        "",
		stream:      stream.New(),
	}
}

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
