package gui

import (
	"bronze/internal/application/services"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateInsertButton(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Inserir Produto", func() {
		w.SetContent(NewInsertWindow(g, w))
	})
}

func CreateListButton(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Lista de Produtos", func() {
		w.SetContent(NewListWindow(g, w))
	})
}

func CreateReportButton(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	statusLabel := widget.NewLabel("")
	return widget.NewButton("Gerar Relatorio", func() {
		err := g.UserService.GenerateReport()
		if err != nil {
			statusLabel.SetText(err.Error())
			return
		}
	})
}

func CreateSaveButton(g *services.GUIService, w fyne.Window, productsList *fyne.Container) (fyne.CanvasObject, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Label) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Produto")

	valueEntry := widget.NewEntry()
	valueEntry.SetPlaceHolder("Valor")

	dataEntry := widget.NewEntry()
	dataEntry.SetPlaceHolder("Mes/Ano")

	statusLabel := widget.NewLabel("")

	return widget.NewButton("Salvar", func() {

		name := nameEntry.Text
		valueStr := valueEntry.Text
		data := dataEntry.Text

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			statusLabel.SetText("Valor invalido!")
			return
		}

		err = g.UserService.SaveProduct(name, data, value)
		if err != nil {
			statusLabel.SetText("Erro ao Inserir!")
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		nameEntry.SetText("")
		valueEntry.SetText("")

		g.RefreshUserList(productsList)
	}), nameEntry, valueEntry, dataEntry, statusLabel
}
