package main

import (
	"errors"

	Fyne "fyne.io/fyne/v2"
	FyneApp "fyne.io/fyne/v2/app"
	FyneContainer "fyne.io/fyne/v2/container"
	FyneDialog "fyne.io/fyne/v2/dialog"
	FyneWidget "fyne.io/fyne/v2/widget"
)

func StartGui() {
	app := FyneApp.New()

	w := app.NewWindow("Časovač | Firesport.cz")
	w.Resize(Fyne.Size{Width: 800, Height: 480})

	hello := FyneWidget.NewLabel("Hello Fyne!")
	hellao := FyneWidget.NewLabel("Hello Fyne!a")
	butt := FyneWidget.NewButton("Hi!", func() {
		hello.SetText("Welcome :)")

		FyneDialog.ShowError(errors.New("toto jest chyba, lebo jsi kokotko"), w)
	})

	w.SetContent(FyneContainer.NewVBox(
		hello,
		hellao,
		butt,
	))

	w.ShowAndRun()
}
