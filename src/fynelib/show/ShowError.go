package show

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/librarygo/src/check"
)

type (
	Interface interface {
		//canvasobjectapi.Interface
		Error()
	}
	object struct{}
)

func Error() { New().Error() }
func (o *object) Error() {
	w := fyne.CurrentApp().NewWindow("Error Show")
	w.SetContent(o.CanvasObject(w))
	w.Show()
}

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	start := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	start.SetText(check.Body())
	return start
}

func New() Interface { return &object{} }
