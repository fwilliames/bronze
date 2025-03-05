package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewReportWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	buttonToInsertWindow := CreateButtonToInsertWindow(g, w)
	dataFilter := CreateSelectFilter(g, w, "data")
	reportButton := CreateReportButton(g, w)
	buttonToListWindow := CreateButtonToListWindow(g, w)

	return container.NewVBox(
		dataFilter,
		reportButton,
		buttonToListWindow,
		buttonToInsertWindow,
	)
}
