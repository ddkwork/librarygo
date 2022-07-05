package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ddkwork/librarygo/src/fynelib/notes"
)

func main() {
	a := app.NewWithID("com.fynelabs.notes")
	//cloud.Enable(a)
	a.SetIcon(notes.ResourceIconPng)
	a.Settings().SetTheme(&notes.MyTheme{})
	w := a.NewWindow("Fyne Notes")

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Sync...", func() {
				//cloud.ShowSettings(a, w)
			}))))

	list := &notes.Notelist{Pref: a.Preferences()}
	list.Load()
	notesUI := &notes.Ui{Notes: list}

	w.SetContent(notesUI.LoadUI())
	notesUI.RegisterKeys(w)

	w.Resize(fyne.NewSize(500, 320))
	w.ShowAndRun()
}
