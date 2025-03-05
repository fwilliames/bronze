package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainWindow(guiService *services.GUIService, w fyne.Window) fyne.CanvasObject {

	buttonToInsertWindow := CreateButtonToInsertWindow(guiService, w)
	buttonToListWindow := CreateButtonToListWindow(guiService, w)
	buttonToReportWindow := CreateButtonToReportWindow(guiService, w)

	return container.NewVBox(
		buttonToInsertWindow,
		buttonToListWindow,
		buttonToReportWindow,
	)
}
