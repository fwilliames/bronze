package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindow(userService *services.UserService, w fyne.Window) fyne.CanvasObject {
	projectLabel := widget.NewLabel("Cadastro de Usuários")

	// Botões principais
	cadastroButton := widget.NewButton("Tela de Cadastro", func() {
		w.SetContent(NewCadastroForm(userService, w))
	})

	listButton := widget.NewButton("Tela de Listagem", func() {
		w.SetContent(NewUserList(userService, w))
	})

	// Retorna o conteúdo da tela principal
	return container.NewVBox(
		projectLabel,
		cadastroButton,
		listButton,
	)
}
