package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewInsertWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	productsList := container.NewVBox()

	buttonSave, nameEntry, valueEntry, dataEntry, statusLabel := CreateSaveButton(g, w, productsList)
	buttonToListWindow := CreateButtonToListWindow(g, w)
	buttonToReportWindow := CreateButtonToReportWindow(g, w)
	marketFilter := CreateSelectFilter(g, w, "market")

	return container.NewVBox(
		nameEntry,
		valueEntry,
		dataEntry,
		marketFilter,
		buttonSave,
		buttonToListWindow,
		buttonToReportWindow,
		statusLabel,
		productsList,
	)
}
