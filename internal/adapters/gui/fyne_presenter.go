package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp(g *services.GUIService) {
	a := app.New()

	w := a.NewWindow("Guia de Supermercado")
	w.Resize(fyne.NewSize(856, 960))
	w.SetContent(NewMainWindow(g, w))
	w.ShowAndRun()
}
