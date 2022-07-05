package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/librarygo/src/check"
	"github.com/ddkwork/librarygo/src/fynelib/fyneTheme"
	"github.com/ddkwork/librarygo/src/fynelib/show"
)

func main() {
	a := app.NewWithID("com.rows.app")
	a.SetIcon(nil)
	fyneTheme.New().SetDarkTheme(a)
	w := a.NewWindow("app")
	w.Resize(fyne.NewSize(640, 480))
	w.SetMaster()
	w.CenterOnScreen()
	w.SetContent(widget.NewButton("测试中文", func() {
		check.Error(111111111)
		show.Error()
	}))
	w.ShowAndRun()
}
