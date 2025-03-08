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
	buttonSaveMarket, marketEntry, statusMarketLabel := CreateSaveMarketButton(guiService, w, marketList)
	buttonSaveData, dataEntry, statusDataLabel := CreateSaveDateButton(guiService, w)

	vbox1 := container.NewVBox(marketEntry, buttonSaveMarket, statusMarketLabel)
	vbox2 := container.NewVBox(dataEntry, buttonSaveData, statusDataLabel)
	hbox := container.NewHBox(buttonToInsertWindow, buttonToListWindow, buttonToReportWindow)

	return container.NewVBox(
		vbox1,
		vbox2,
		hbox,
	)
}
