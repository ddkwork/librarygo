package myc2go

import (
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
		defines     []string
		apis        []string
		replaces    map[string]string
		files       []string
		skip        string
	}
)

func NewSetup(setup Setup) Interface {
	return &object{
		contains:    setup.SetContains(),
		containsNot: setup.SetContainsNot(),
		root:        setup.SetRoot(),
		ext:         setup.SetExt(),
		defines:     nil,
		apis:        nil,
		replaces:    nil,
		files:       nil,
		skip:        setup.SetSkip(),
	}
}

func (o *object) ConvertAll() (ok bool) {
	s := stream.New()
	for _, p := range o.root {
		if !mycheck.Error(filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {

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
	for _, file := range o.files {
		readFile, err := os.ReadFile(file)
		if !mycheck.Error(err) {
			return
		}
		source, err := format.Source(readFile)
		if !mycheck.Error(err) {
			return
		}
		if !tool.File().WriteTruncate(file, source) {
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

func New() Interface { return &object{} }

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
