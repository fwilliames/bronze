package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewListWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {

	productsList := container.NewVBox()

	buttonToInsertWindow := CreateButtonToInsertWindow(g, w)
	buttonToReportWindow := CreateButtonToReportWindow(g, w)
	buttonToMainWindow := CreateButtonToMainWindow(g, w)

	dataSelectFilterAndList := CreateDataSelectFilter(g, w, productsList)

	vbox := container.NewVBox(dataSelectFilterAndList, productsList)
	hbox := container.NewHBox(buttonToInsertWindow, buttonToReportWindow, buttonToMainWindow)

	return container.NewVBox(
		vbox,
		hbox,
	)
}
