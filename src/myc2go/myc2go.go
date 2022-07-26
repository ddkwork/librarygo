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
		SetBasicTypes(basicTypes string)
	}
	Setup struct {
		SetRoot        func() []string
		SetExt         func() []string
		SetContains    func() []string
		SetContainsNot func() []string
		SetSkip        func() string
		SetReplaces    func() map[string]string
		SetBasicTypes  func() string
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
		basicTypes  string
	}
)

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
func NewSetup(setup Setup) Interface {
	return &object{
		contains:    setup.SetContains(),
		containsNot: setup.SetContainsNot(),
		root:        setup.SetRoot(),
		ext:         setup.SetExt(),
		apis:        nil,
		replaces:    setup.SetReplaces(),
		files:       make(map[string]string), //goFilePath goFileBody
		skip:        setup.SetSkip(),
		stream:      stream.New(),
		basicTypes:  setup.SetBasicTypes(),
	}
}

func (o *object) HandleBasicTypes(path string) (ok bool) {
	base := filepath.Base(filepath.Dir(path))
	if base == o.basicTypes {
		cFiles := make([]string, 0)
		if !mycheck.Error(filepath.Walk(filepath.Dir(path), func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == `.h` {
				abs, err := filepath.Abs(path)
				if !mycheck.Error(err) {
					return err
				}
				c := strings.ReplaceAll(abs, `.h`, `.c`)
				cFiles = append(cFiles, c)
			}
			return err
		})) {
			return
		}

		for _, file := range cFiles {
			s := stream.NewString(BasicTypes)
			b, err := os.ReadFile(path)
			if !mycheck.Error(err) {
				return
			}
			all := strings.ReplaceAll(string(b), "#pragma once", "")
			s.WriteStringLn(all)
			if !tool.File().WriteTruncate(file, s.String()) {
				return
			}
		}
		for _, file := range cFiles {
			abs, err := filepath.Abs(file)
			if !mycheck.Error(err) {
				return
			}
			Run(abs, o.basicTypes)
		}

		goFiles := make([]string, 0)

		if !mycheck.Error(filepath.Walk(filepath.Dir(path), func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == `.go` {
				goFiles = append(goFiles, path)
			}
			return err
		})) {
			return
		}
		for _, file := range goFiles {
			b, err := os.ReadFile(file)
			if !mycheck.Error(err) {
				return
			}
			all := strings.ReplaceAll(string(b), cType, ``)
			if !tool.File().WriteTruncate(file, all) {
				return
			}
		}
		if !tool.File().WriteTruncate(filepath.Join(filepath.Dir(path), "cType.go"), cType) {
			return
		}
	}
	return true
}

const cType = `
type DWORD = uint32
type BOOL = int32
type BYTE = uint8
type WORD = uint16
type FLOAT = float32
type PFLOAT = *float32
type INT = int32
type UINT = uint32
type PUINT = *uint32
type PBOOL = *int32
type LPBOOL = *int32
type PBYTE = *uint8
type LPBYTE = *uint8
type PINT = *int32
type LPINT = *int32
type PWORD = *uint16
type LPWORD = *uint16
type LPLONG = *int32
type PDWORD = *uint32
type LPDWORD = *uint32
type LPVOID = unsafe.Pointer
type PVOID = unsafe.Pointer
type LPCVOID = unsafe.Pointer
type ULONG = uint32
type PULONG = *uint32
type USHORT = uint16
type PUSHORT = *uint16
type UCHAR = uint8
type PUCHAR = *uint8
type CHAR = int8
type SHORT = int16
type LONG = int32
type QWORD = uint64
type UINT64 = uint64
type PUINT64 = *uint64
type ULONG64 = uint64
type PULONG64 = *uint64
type DWORD64 = uint64
type PDWORD64 = *uint64
type BOOLEAN = uint8
type PBOOLEAN = *uint8
type INT8 = int8
type PINT8 = *int8
type INT16 = int16
type PINT16 = *int16
type INT32 = int32
type PINT32 = *int32
type INT64 = int64
type PINT64 = *int64
type UINT8 = uint8
type PUINT8 = *uint8
type UINT16 = uint16
type PUINT16 = *uint16
type UINT32 = uint32
type PUINT32 = *uint32
type wchar_t = *int16
type WCHAR = *int16
`

func (o *object) HandleDefines(path string) (ok bool) {
	body, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	lines, ok := tool.File().ToLines(body)
	if !ok {
		return
	}
	ss := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, o.skip) {
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
		for orig, replace := range o.replaces {
			if strings.Contains(define, orig) {
				define = strings.ReplaceAll(define, orig, replace)
			}
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
	b, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	body, ok := tool.File().ToLines(b)
	if !ok {
		return
	}
	apis := make([]string, 0)
	for i, line := range body {
		if strings.Contains(line, "funcs") { //todo
			apis = apis[i+1:]
			break
		}
	}
	lines, ok2 := tool.File().ToLines(o.stream.String())
	if !ok2 {
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
									if !o.HandleBasicTypes(path) {
										return err
									}
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
									//if !o.HandleApis(path) {
									//	return err
									//}
									if !o.HandleDefines(path) {
										return err
									}
									o.files[goFilePath] = o.stream.String()
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
	//for _, file := range o.files {
	//	Run(file)
	//}
	return o.Format()
}
func (o *object) Format() (ok bool) {
	for path, body := range o.files {
		if !tool.File().WriteTruncate(path, body) {
			return
		}
	}
	//for path := range o.files {//todo move to last
	//	if !cmd.Run(`gofmt -l -w ` + path) {
	//		return
	//	}
	//}
	return true
}
func (o *object) SetSkip(skip string)                    { o.skip = skip }
func (o *object) SetExt(ext []string)                    { o.ext = ext }
func (o *object) SetRoot(root []string)                  { o.root = root }
func (o *object) SetContainsNot(containsNot []string)    { o.containsNot = containsNot }
func (o *object) SetContains(contains []string)          { o.contains = contains }
func (o *object) SetReplaces(replaces map[string]string) { o.replaces = replaces }
func (o *object) SetBasicTypes(basicTypes string)        { o.basicTypes = basicTypes }

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

const BasicTypes = `
#define MAX_PATH          260
typedef unsigned long DWORD;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef float FLOAT;
typedef FLOAT *PFLOAT;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef float FLOAT;
typedef FLOAT *PFLOAT;
typedef BOOL *PBOOL;
typedef BOOL *LPBOOL;
typedef BYTE *PBYTE;
typedef BYTE *LPBYTE;
typedef int *PINT;
typedef int *LPINT;
typedef WORD *PWORD;
typedef WORD *LPWORD;
typedef long *LPLONG;
typedef DWORD *PDWORD;
typedef DWORD *LPDWORD;
typedef void *LPVOID;
typedef void *PVOID;
typedef const void *LPCVOID;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef unsigned long ULONG;
typedef ULONG *PULONG;
typedef unsigned short USHORT;
typedef USHORT *PUSHORT;
typedef unsigned char UCHAR;
typedef UCHAR *PUCHAR;
typedef char CHAR;
typedef short SHORT;
typedef long LONG;
typedef int INT;


typedef unsigned long long QWORD;
typedef unsigned __int64   UINT64, *PUINT64;
typedef unsigned long      DWORD;
typedef int                BOOL;
typedef unsigned char      BYTE;
typedef unsigned short     WORD;
typedef int                INT;
typedef unsigned int       UINT;
typedef unsigned int *     PUINT;
typedef unsigned __int64   ULONG64, *PULONG64;
typedef unsigned __int64   DWORD64, *PDWORD64;
typedef char               CHAR;

typedef void *LPVOID;
typedef void *PVOID;
typedef const void *LPCVOID;

typedef unsigned char    UCHAR;
typedef unsigned short   USHORT;
typedef unsigned long    ULONG;
typedef UCHAR            BOOLEAN;  // winnt
typedef BOOLEAN *        PBOOLEAN; // winnt
typedef signed char      INT8, *PINT8;
typedef signed short     INT16, *PINT16;
typedef signed int       INT32, *PINT32;
typedef signed __int64   INT64, *PINT64;
typedef unsigned char    UINT8, *PUINT8;
typedef unsigned short   UINT16, *PUINT16;
typedef unsigned int     UINT32, *PUINT32;
typedef unsigned __int64 UINT64, *PUINT64;

typedef             PINT16 wchar_t;
typedef wchar_t            WCHAR;
`
