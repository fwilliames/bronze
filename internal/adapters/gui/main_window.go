package gui

import (
	"bronze/internal/application/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainWindow(guiService *services.GUIService, w fyne.Window) fyne.CanvasObject {

	marketList := container.NewVBox()

	buttonToInsertWindow := CreateButtonToInsertWindow(guiService, w)
	buttonToListWindow := CreateButtonToListWindow(guiService, w)
	buttonToReportWindow := CreateButtonToReportWindow(guiService, w)
	buttonSaveMarket, marketEntry, statusLabel := CreateSaveMarketButton(guiService, w, marketList)

	return container.NewVBox(
		statusLabel,
		buttonToInsertWindow,
		buttonToListWindow,
		buttonToReportWindow,
		marketEntry,
		buttonSaveMarket,
	)
}
