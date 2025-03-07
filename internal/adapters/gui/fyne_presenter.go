package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2/app"
)

func StartApp(g *services.GUIService) {
	a := app.New()

	w := a.NewWindow("Guia de Supermercado")
	w.SetContent(NewMainWindow(g, w))
	w.CenterOnScreen()
	w.ShowAndRun()

}
