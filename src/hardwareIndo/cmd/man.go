package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/ddkwork/librarygo/src/hardwareIndo"
	"github.com/ddkwork/librarygo/src/hardwareIndo/cmd/hardinfo"
)

func main() {
	a := app.NewWithID("com.rows.app")
	//a.SetIcon(nil)
	//fyneTheme.New().SetDarkTheme(a)
	w := a.NewWindow("hardInfo")
	//w.Resize(fyne.NewSize(140, 580))
	//w.SetMaster()
	w.CenterOnScreen()
	h := hardinfo.New()
	w.SetContent(h.CanvasObject(w))
	w.ShowAndRun()
}

func main1() {
	h := hardwareIndo.New()
	if !h.SsdInfo.Get() { //todo bug cpu pkg init
		return
	}
	if !h.CpuInfo.Get() {
		return
	}
	if !h.MacInfo.Get() {
		return
	}
}
