package gui

import (
	"bronze/internal/application/services"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewCadastroForm retorna o conteúdo da tela de cadastro
func NewCadastroForm(userService *services.UserService, w fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Digite seu nome")

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("Digite sua idade")

	statusLabel := widget.NewLabel("")

	saveButton := widget.NewButton("Salvar", func() {
		name := nameEntry.Text
		ageStr := ageEntry.Text

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			statusLabel.SetText("Idade inválida!")
			return
		}

		err = userService.SaveUser(name, age)
		if err != nil {
			statusLabel.SetText("Erro ao salvar!")
			return
		}

		statusLabel.SetText("Salvo com sucesso!")
		nameEntry.SetText("")
		ageEntry.SetText("")
	})

	listButton := widget.NewButton("Tela de Listagem", func() {
		w.SetContent(NewUserList(userService, w))
	})

	// Voltar para a tela principal
	backButton := widget.NewButton("Voltar", func() {
		w.SetContent(NewMainWindow(userService, w))
	})

	// Retorna o conteúdo da tela de cadastro
	return container.NewVBox(
		widget.NewLabel("Cadastro de Usuários"),
		nameEntry,
		ageEntry,
		saveButton,
		statusLabel,
		listButton,
		backButton,
	)
}
