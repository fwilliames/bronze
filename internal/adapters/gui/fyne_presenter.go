package gui

import (
	"bronze/internal/application/services"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartApp(userService *services.UserService) {

	a := app.New()
	w := a.NewWindow("Cadastro de Usuários")
	w.Resize(fyne.NewSize(300, 200))

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

		statusLabel.SetText(fmt.Sprintf("Salvo: %s, %d anos", name, age))
		nameEntry.SetText("")
		ageEntry.SetText("")
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Cadastro de Usuários"),
		nameEntry,
		ageEntry,
		saveButton,
		statusLabel,
	))

	w.ShowAndRun()
}
