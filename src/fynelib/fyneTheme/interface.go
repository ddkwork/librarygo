package fyneTheme

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed ttf/HarmonyOS_Sans_SC_Light.ttf
var ttfBuf []byte

var HarmonyOS_Sans_SC_Bold_Ttf = &fyne.StaticResource{
	StaticName:    "HarmonyOS_Sans_SC_Bold.ttf",
	StaticContent: ttfBuf,
}

type (
	myTheme interface {
		SetDarkTheme(a fyne.App)
		SetLightTheme(a fyne.App)
	}
	_theme struct {
		*darkTheme
		*lightTheme
	}
)

func New() myTheme {
	return &_theme{
		darkTheme:  &darkTheme{},
		lightTheme: &lightTheme{},
	}
}

func SetDarkTheme(a fyne.App)  { New().SetDarkTheme(a) }
func SetLightTheme(a fyne.App) { New().SetLightTheme(a) }

func (p *_theme) SetDarkTheme(a fyne.App) {
	a.Settings().SetTheme(p.darkTheme)
}
func (p *_theme) SetLightTheme(a fyne.App) {
	a.Settings().SetTheme(p.lightTheme)
}
