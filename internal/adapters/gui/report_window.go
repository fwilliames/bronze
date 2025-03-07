package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewReportWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {

	buttonToInsertWindow := CreateButtonToInsertWindow(g, w)
	dataFilter := CreateSelectFilter(g, w, "data")
	marketFilter := CreateSelectFilter(g, w, "market")
	reportButton := CreateReportButton(g, w)
	buttonToListWindow := CreateButtonToListWindow(g, w)
	buttonToMainWindow := CreateButtonToMainWindow(g, w)

	vbox := container.NewVBox(dataFilter, marketFilter, reportButton)
	hbox := container.NewHBox(buttonToInsertWindow, buttonToListWindow, buttonToMainWindow)

	return container.NewVBox(
		vbox,
		hbox,
	)
}
