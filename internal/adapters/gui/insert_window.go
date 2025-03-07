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
	buttonToMainWindow := CreateButtonToMainWindow(g, w)
	marketFilter := CreateSelectFilter(g, w, "market")

	vbox := container.NewVBox(nameEntry, valueEntry, dataEntry, marketFilter, buttonSave, statusLabel, productsList)
	hbox := container.NewHBox(buttonToListWindow, buttonToReportWindow, buttonToMainWindow)

	return container.NewVBox(
		vbox,
		hbox,
	)
}
