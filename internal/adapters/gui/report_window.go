package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewReportWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	//marketsList := container.NewVBox()
	buttonToInsertWindow := CreateButtonToInsertWindow(g, w)
	dataFilter := CreateSelectFilter(g, w, "data")
	marketFilter := CreateSelectFilter(g, w, "market")
	reportButton := CreateReportButton(g, w)
	buttonToListWindow := CreateButtonToListWindow(g, w)

	return container.NewVBox(
		dataFilter,
		marketFilter,
		reportButton,
		buttonToListWindow,
		buttonToInsertWindow,
	)
}
