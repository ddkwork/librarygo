package driverTool

import "fyne.io/fyne/v2/widget"

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	return widget.NewButton("todo", func() {})
}

func New() Interface { return &object{} }
