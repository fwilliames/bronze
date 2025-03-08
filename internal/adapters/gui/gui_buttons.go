package gui

import (
	"bronze/internal/application/services"
	"bronze/internal/config/utils"
	"log"
	"strconv"
	"strings"

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

func CreateButtonToMainWindow(g *services.GUIService, w fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Inicio", func() {
		w.SetContent(NewMainWindow(g, w))
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

	quantityEntry := widget.NewEntry()
	quantityEntry.SetPlaceHolder("Quantidade")

	dataEntry := widget.NewEntry()
	dataEntry.SetPlaceHolder("Mes/Ano")

	statusLabel := widget.NewLabel("")

	return widget.NewButton("Salvar", func() {

		name := nameEntry.Text
		valueStr := valueEntry.Text
		quantityStr := quantityEntry.Text

		valueStr = strings.TrimSpace(valueStr) // Remove espaços extras
		if valueStr == "" {
			statusLabel.SetText("O valor não pode estar vazio!")
			return
		}
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			statusLabel.SetText("Valor invalido!")
			return
		}

		quantityStr = strings.TrimSpace(quantityStr) // Remove espaços extras

		if quantityStr == "" {
			statusLabel.SetText("O valor não pode estar vazio!")
			return
		}

		quantity, err := strconv.ParseInt(quantityStr, 10, 64)
		if err != nil {
			statusLabel.SetText("Valor inválido! Digite um número inteiro.")
			return
		}

		totalValue := utils.Prod(value, quantity)

		err = g.UserService.SaveProduct(name, g.Filters.Data, g.Filters.Market, value, totalValue, quantity)
		if err != nil {
			statusLabel.SetText(err.Error())
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		nameEntry.SetText("")
		valueEntry.SetText("")

		g.ListProductsByFilter(productsList, g.Filters.Data)
	}), nameEntry, valueEntry, quantityEntry, statusLabel
}

func CreateDataSelectFilter(g *services.GUIService, w fyne.Window, productsList *fyne.Container) *widget.Select {
	dates, err := g.UserService.GetDates()
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

	switch field {

	case "data":

		datas, err := g.UserService.GetDates()
		if err != nil {
			log.Println("Erro ao recuperar datas:", err)
		}

		selectWidget := widget.NewSelect(datas, func(selected string) {
			g.Filters.Data = selected
		})

		selectWidget.PlaceHolder = "Selecione uma Data"

		return selectWidget

	case "market":
		markets, err := g.UserService.GetMarkets()
		if err != nil {
			log.Println("Erro ao recuperar Markets:", err)
			markets = []string{"Nenhuma data encontrada"}
		}

		selectWidget := widget.NewSelect(markets, func(selected string) {
			g.Filters.Market = selected
		})

		selectWidget.PlaceHolder = "Selecione um Mercado"

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

func CreateSaveDateButton(g *services.GUIService, w fyne.Window) (fyne.CanvasObject, *widget.Entry, *widget.Label) {

	dataEntry := widget.NewEntry()
	dataEntry.SetPlaceHolder("Data")

	statusLabel := widget.NewLabel("")

	return widget.NewButton("Salvar Mes/Ano", func() {

		err := g.UserService.SaveData(dataEntry.Text)
		if err != nil {
			statusLabel.SetText("Erro ao Inserir!")
			return
		}

		statusLabel.SetText("Inserido com sucesso!")
		dataEntry.SetText("")
		dataEntry.SetPlaceHolder("Mercado")

	}), dataEntry, statusLabel
}
