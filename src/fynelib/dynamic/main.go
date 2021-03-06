package main

import (
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/fpabl0/sparky-go"
	"github.com/fpabl0/sparky-go/scont"
	"github.com/fpabl0/sparky-go/svalid"
	"github.com/fpabl0/sparky-go/swid"
)

const (
	keyFirstName sparky.ValueKey = iota
)

func main() {
	a := app.New()
	// a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("hello")
	w.Resize(fyne.NewSize(500, 400))

	svalid.ConfigErrMessages(&svalid.ErrorMessages{
		NotEmpty: "* Campo requerido",
	})

	ctx := sparky.NewContext(w)
	ctx.PutValue(keyFirstName, "Pablo")

	l := ctx.ShowLoader("Loading...")
	go func() {
		time.Sleep(1 * time.Second)
		<-l.Error("Un simple error muy pero muy largo debe de salir mal ahora!!!!!!")
		ctx.ShowInfo("Atención!", "Hay una nueva versión disponible")
		firstName := ctx.GetValue(keyFirstName).(string)
		fmt.Printf("Listo %s\n", firstName)
	}()

	// w.SetContent(content(ctx))

	notEmptyField := func(n int) *swid.TextFormField {
		forbidden := "wrong"
		if n == 2 {
			forbidden = ""
		}

		tf := swid.NewTextFormField(fmt.Sprintf("Name %d", n), "")
		tf.Validator = func(s string) error {
			if s == forbidden {
				return errors.New("wrong")
			}
			return nil
		}
		tf.OnSaved = func(s string) { fmt.Println("saved:", s) }

		tf.Placeholder = "Write your name"
		tf.Hint = "A hint text"
		if n == 1 {
			tf.SetText("wron")
			tf.Disable()
		}
		return tf
	}

	pf := swid.NewPasswordTextFormField("Password", "1234")
	pf.Hint = "Your password"

	mf := swid.NewMaskedTextFormField("Masked", "", "99/99", "99/99")
	mf.Hint = "A masked form field"
	mf.ActionItem = widget.NewIcon(theme.ComputerIcon())

	sef := swid.NewSelectEntryFormField("OS", "", []string{"Mac", "Windows"})
	sef.Validator = svalid.NotEmpty()
	sef.Hint = "Your favorite OS"

	sf := swid.NewSelectFormField("Car", "", []string{"Audi", "Toyota"})
	sf.OnSaved = func(s string) { fmt.Println(s) }
	sf.Hint = "Your car"

	f := swid.NewForm(2,
		notEmptyField(1), notEmptyField(2),
		notEmptyField(3), notEmptyField(4),
		pf, mf,
		sef, sf,
	)

	f.OnValidationChanged = func(v bool) {
		fmt.Println("valid: ", v)
	}

	tf := swid.NewTextField()
	tf.MaxLength = 4

	e := widget.NewEntry()
	e.Text = "wrong"
	e.Validator = func(s string) error {
		if s == "wrong" {
			return errors.New("wrong")
		}
		return nil
	}
	e.Validate()
	e.Disable()

	sl := swid.NewSelectFormField("Name", "", []string{"Peter", "Ana"})
	sl.Disable()

	w.SetContent(container.NewVBox(
		f,
		swid.NewMaskedTextField("+(999) 999-9999", "+(999) 999-9999"),
		swid.NewMaskedTextField("99/99/99", "dd/MM/yy"),
		swid.NewRestrictTextField(swid.RestrictInputInteger),
		swid.NewRestrictTextField(swid.RestrictInputFloat),
		swid.NewRestrictTextField(swid.RestrictInputEmail),
		tf,
		e,
		sl,
		container.NewHBox(
			f.CreateSubmitButton("Crear", func() { f.Save() }),
			f.CreateResetButton("Reset"),
		),
	))

	w.ShowAndRun()
}

func content(ctx sparky.Context) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return 10
		},
		func() fyne.CanvasObject {
			return newTitleItem()
		},
		func(index widget.ListItemID, obj fyne.CanvasObject) {
			tile := obj.(*tileItem)
			tile.Name = fmt.Sprintf("Name %d", index)
			tile.LastName = fmt.Sprintf("LastName %d", index)
			tile.Cellphone = "Cellphone"
			tile.Refresh()
		},
	)

	return scont.NewFrame(2, 2,
		container.NewBorder(widget.NewButton("Press", func() {
			go func() {
				resp := <-ctx.ShowPasswordInput("Nueva contraseña", "Ingrese su nueva contraseña", "Confirmar")
				if resp == nil {
					return
				}
				fmt.Println(*resp)
			}()
		}), nil, nil, nil, list),
	)
}

// ===============================================================
// Title Form
// ===============================================================

type tileItem struct {
	widget.BaseWidget
	Name      string
	LastName  string
	Cellphone string
}

func newTitleItem() *tileItem {
	p := &tileItem{}
	p.ExtendBaseWidget(p)
	return p
}

func (p *tileItem) CreateRenderer() fyne.WidgetRenderer {
	p.ExtendBaseWidget(p)
	return sparky.CreateRenderer(&titleItemRenderer{widget: p})
}

type titleItemRenderer struct {
	name      *canvas.Text
	lastname  *canvas.Text
	cellphone *canvas.Text

	widget *tileItem
}

func (r *titleItemRenderer) CreateContent() *fyne.Container {
	r.name = canvas.NewText(r.widget.Name, theme.ForegroundColor())
	r.name.TextSize = theme.CaptionTextSize()
	r.lastname = canvas.NewText(r.widget.LastName, theme.ForegroundColor())
	r.lastname.TextSize = theme.CaptionTextSize()
	r.cellphone = canvas.NewText(r.widget.Cellphone, theme.ForegroundColor())
	r.cellphone.TextSize = theme.CaptionTextSize()

	return container.NewVBox(
		r.name,
		r.lastname,
		r.cellphone,
	)
}

func (r *titleItemRenderer) Destroy() {}

func (r *titleItemRenderer) Refresh() {
	r.name.Text = r.widget.Name
	r.name.Refresh()

	r.lastname.Text = r.widget.LastName
	r.lastname.Refresh()

	r.cellphone.Text = r.widget.Cellphone
	r.cellphone.Refresh()
}
