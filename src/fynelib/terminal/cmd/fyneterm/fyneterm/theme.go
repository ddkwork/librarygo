package fyneterm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type termTheme struct {
	fyne.Theme
}

func NewTermTheme() fyne.Theme {
	return &termTheme{fyne.CurrentApp().Settings().Theme()}
}

// Color fixes a bug < 2.1 where theme.DarkTheme() would not override user preference.
func (t *termTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case termOverlay:
		if c := t.Color("fynedeskPanelBackground", v); c != color.Transparent {
			return c
		}
		if v == theme.VariantLight {
			return color.NRGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xf6}
		}
		return color.NRGBA{R: 0x16, G: 0x16, B: 0x16, A: 0xf6}
	}
	return t.Theme.Color(n, theme.VariantDark)
}

func (t *termTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(n)
}
