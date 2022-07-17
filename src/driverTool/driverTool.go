package driverTool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/librarygo/src/driverTool/driver"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/fpabl0/sparky-go/swid"
	"io/fs"
	"path/filepath"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{ drivers []string }
)

func New() Interface { return &object{drivers: make([]string, 0)} }

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	o.WalkAllDriverPath("D:\\codespace\\workspace\\src\\cppkit\\driverTool")

	logView := widget.NewMultiLineEntry()
	logView.PlaceHolder = "log ..."

	path := swid.NewSelectFormField("path", "", o.drivers)
	link := swid.NewTextFormField("link", "")
	path.OnChanged = func(s string) {
		ext := filepath.Ext(s)
		base := filepath.Base(s)
		base = base[:len(base)-len(ext)]
		link.SetText(base)
	}
	ioCode := swid.NewTextFormField("ioCode", "")
	d := driver.New()
	load := widget.NewButton("load", func() {
		if !d.Load(path.Selected()) {
			logView.SetText(mycheck.Body())
			return
		}
		logView.SetText("load " + path.Selected() + " successful")
	})
	unload := widget.NewButton("unload", func() {
		if !d.Unload() {
			logView.SetText(mycheck.Body())
			return
		}
		logView.SetText("unload " + path.Selected() + " successful")
	})
	errCode := swid.NewTextFormField("errCode", "")
	ntstatus := swid.NewTextFormField("ntstatus", "")
	hresult := swid.NewTextFormField("hresult", "")
	winerror := swid.NewTextFormField("winerror", "")

	reload := swid.NewTextFormField("reload path", "")
	reload.OnChanged = func(s string) {
		o.drivers = o.drivers[:0]
		path.Options = path.Options[:0]
		o.WalkAllDriverPath(s)
		path.Options = o.drivers
	}
	form := container.NewGridWithColumns(1,
		reload,
		path,
		link,
		ioCode,
		errCode,
		ntstatus,
		hresult,
		winerror,
		container.NewGridWithColumns(2, load, unload),
	)
	split := container.NewHSplit(form, logView)
	split.Offset = 0.4
	return split
}

func (o *object) WalkAllDriverPath(root string) bool {
	newRoot := root
	if root == "" {
		newRoot = "."
	}
	i := 0
	return mycheck.Error(filepath.Walk(newRoot, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".sys" {
			i++
			o.drivers = append(o.drivers, path)
		}
		return nil
	}))
}
