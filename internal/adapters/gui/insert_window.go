package gui

import (
	"bronze/internal/application/services"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewCadastroForm retorna o conteúdo da tela de cadastro
func NewCadastroForm(guiService *services.GUIService, w fyne.Window) fyne.CanvasObject {

	listContainer := container.NewVBox()

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Produto")

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("Valor")

	statusLabel := widget.NewLabel("")

	saveButton := widget.NewButton("Inserir", func() {
		name := nameEntry.Text
		ageStr := ageEntry.Text

		age, err := strconv.ParseFloat(ageStr, 64)
		if err != nil {
			statusLabel.SetText("Valor invalido!")
			return
		}

		err = guiService.UserService.SaveUser(name, age)
		if err != nil {
			statusLabel.SetText("Erro ao Inserir!")
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		nameEntry.SetText("")
		ageEntry.SetText("")

		guiService.RefreshUserList(guiService.UserService, listContainer)
	})

	listButton := widget.NewButton("Lista de Produtos", func() {
		w.SetContent(NewUserList(guiService, w))
	})

	// Voltar para a tela principal
	reportButton := widget.NewButton("Gerar Relatorio", func() {
	})

	guiService.RefreshUserList(guiService.UserService, listContainer)

	// Retorna o conteúdo da tela de cadastro
	return container.NewVBox(
		nameEntry,
		ageEntry,
		saveButton,
		listButton,
		reportButton,
		statusLabel,
		listContainer,
	)
}
