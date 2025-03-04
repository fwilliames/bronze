package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp(userService *services.UserService) {
	a := app.New()
	w := a.NewWindow("Guia de Supermercado")
	w.Resize(fyne.NewSize(856, 960))

	// Exibe os componentes principais (formulário de cadastro e botões)
	w.SetContent(NewMainWindow(userService, w))

	w.ShowAndRun()
}
