package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewInsertWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	productsList := container.NewVBox()

	saveButton, nameEntry, valueEntry, statusLabel := CreateSaveButton(g, w, productsList)
	listButton := CreateListButton(g, w)
	reportButton := CreateReportButton(g, w)

	g.RefreshUserList(productsList)

	return container.NewVBox(
		nameEntry,
		valueEntry,
		saveButton,
		listButton,
		reportButton,
		statusLabel,
		productsList,
	)
}
