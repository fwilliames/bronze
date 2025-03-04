package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindow(userService *services.UserService, w fyne.Window) fyne.CanvasObject {
	// Botões principais
	insertButton := widget.NewButton("Inserir Produto", func() {
		w.SetContent(NewCadastroForm(userService, w))
	})

	listButton := widget.NewButton("Lista de Produtos", func() {
		w.SetContent(NewUserList(userService, w))
	})

	reportButton := widget.NewButton("Gerar Relatorio", func() {
		w.SetContent(NewUserList(userService, w))
	})

	// Retorna o conteúdo da tela principal
	return container.NewVBox(
		insertButton,
		listButton,
		reportButton,
	)
}
