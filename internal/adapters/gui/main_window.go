package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainWindow(guiService *services.GUIService, w fyne.Window) fyne.CanvasObject {

	insertButton := CreateInsertButton(guiService, w)
	listButton := CreateListButton(guiService, w)
	reportButton := CreateReportButton(guiService, w)

	return container.NewVBox(
		insertButton,
		listButton,
		reportButton,
	)
}
