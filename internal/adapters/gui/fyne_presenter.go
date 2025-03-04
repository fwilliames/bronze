package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp(userService *services.UserService) {
	a := app.New()
	w := a.NewWindow("Cadastro de Usuários")
	w.Resize(fyne.NewSize(300, 200))

	// Exibe os componentes principais (formulário de cadastro e botões)
	w.SetContent(NewMainWindow(userService, w))

	w.ShowAndRun()
}
