package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindow(guiService *services.GUIService, w fyne.Window) fyne.CanvasObject {
	// Botões principais
	insertButton := widget.NewButton("Inserir Produto", func() {
		w.SetContent(NewCadastroForm(guiService, w))
	})

	listButton := widget.NewButton("Lista de Produtos", func() {
		w.SetContent(NewUserList(guiService, w))
	})

	reportButton := widget.NewButton("Gerar Relatorio", func() {
		w.SetContent(NewUserList(guiService, w))
	})

	// Retorna o conteúdo da tela principal
	return container.NewVBox(
		insertButton,
		listButton,
		reportButton,
	)
}
