package gui

import (
	"bronze/internal/application/services"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateButtonToInsertWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Inserir Produto", func() {
		w.SetContent(NewInsertWindow(g, w))
	})
}

func CreateButtonToListWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Lista de Produtos", func() {
		w.SetContent(NewListWindow(g, w))
	})
}

func CreateButtonToReportWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Relatorios", func() {
		w.SetContent(NewReportWindow(g, w))
	})
}

func CreateReportButton(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	statusLabel := widget.NewLabel("")
	return widget.NewButton("Gerar Relatorio", func() {
		err := g.UserService.GenerateReport(g.Filters)
		if err != nil {
			statusLabel.SetText(err.Error())
			return
		}
	})
}

func CreateSaveButton(g *services.GUIService, w fyne.Window, productsList *fyne.Container) (fyne.CanvasObject, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Label) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Produto")

	valueEntry := widget.NewEntry()
	valueEntry.SetPlaceHolder("Valor")

	dataEntry := widget.NewEntry()
	dataEntry.SetPlaceHolder("Mes/Ano")

	statusLabel := widget.NewLabel("")

	return widget.NewButton("Salvar", func() {

		name := nameEntry.Text
		valueStr := valueEntry.Text
		data := dataEntry.Text

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			statusLabel.SetText("Valor invalido!")
			return
		}

		err = g.UserService.SaveProduct(name, data, g.Filters.Market, value)
		if err != nil {
			statusLabel.SetText("Erro ao Inserir!")
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		nameEntry.SetText("")
		valueEntry.SetText("")

		g.ListProductsByFilter(productsList, data)
	}), nameEntry, valueEntry, dataEntry, statusLabel
}

func CreateDataSelectFilter(g *services.GUIService, w fyne.Window, productsList *fyne.Container) *widget.Select {
	dates, err := g.UserService.GetUniqueDates()
	if err != nil {
		log.Println("Erro ao recuperar datas:", err)
		dates = []string{"Nenhuma data encontrada"}
	}

	selectWidget := widget.NewSelect(dates, func(selected string) {
		g.ListProductsByFilter(productsList, selected)
	})

	selectWidget.PlaceHolder = "Selecione uma data"

	return selectWidget
}

func CreateSelectFilter(g *services.GUIService, w fyne.Window, field string) *widget.Select {

	var items []string

	switch field {

	case "data":
		datas, err := g.UserService.GetUniqueDates()
		if err != nil {
			log.Println("Erro ao recuperar datas:", err)
			datas = []string{"Nenhuma data encontrada"}
		}
		items = datas
		selectWidget := widget.NewSelect(items, func(selected string) {
			g.Filters.Data = selected
		})

		selectWidget.PlaceHolder = "Selecione um item"

		return selectWidget

	case "market":
		markets, err := g.UserService.GetMarkets()
		if err != nil {
			log.Println("Erro ao recuperar Markets:", err)
			markets = []string{"Nenhuma data encontrada"}
		}
		items = markets
		selectWidget := widget.NewSelect(items, func(selected string) {
			g.Filters.Market = selected
		})

		selectWidget.PlaceHolder = "Selecione um item"

		return selectWidget
	}
	return nil
}

func CreateSaveMarketButton(g *services.GUIService, w fyne.Window, marketList *fyne.Container) (fyne.CanvasObject, *widget.Entry, *widget.Label) {

	marketEntry := widget.NewEntry()
	marketEntry.SetPlaceHolder("Mercado")

	statusLabel := widget.NewLabel("")

	return widget.NewButton("Salvar Mercado", func() {

		err := g.UserService.SaveMarket(marketEntry.Text)
		if err != nil {
			statusLabel.SetText("Erro ao Inserir!")
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		marketEntry.SetText("")
		marketEntry.SetPlaceHolder("Mercado")

		g.ListProductsByFilter(marketList, "SuperMercados")
	}), marketEntry, statusLabel
}
