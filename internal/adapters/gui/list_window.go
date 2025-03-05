package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewListWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {

	productsList := container.NewVBox()
	insertButton := CreateInsertButton(g, w)
	reportButton := CreateReportButton(g, w)
	dataSelectFilter := CreateDataSelectFilter(g, w, productsList)

	return container.NewVBox(
		dataSelectFilter,
		productsList,
		insertButton,
		reportButton,
	)
}
