package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("My Fyne App")
	w.SetContent(widget.NewLabel("Basic Fyne app created."))
	w.ShowAndRun()
}
