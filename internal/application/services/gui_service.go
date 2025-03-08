package services

import (
	"bronze/internal/config/colors"
	"bronze/internal/config/utils"
	"bronze/internal/domain"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GUIService struct {
	UserService *UserService
	Filters     Filters
}

type Filters struct {
	Data   string
	Market string
}

func NewGUIService(userService *UserService) *GUIService {
	return &GUIService{UserService: userService}
}

func createListToShow(titles ...string) *fyne.Container {

	headers := make([]*canvas.Text, len(titles))
	background := canvas.NewRectangle(color.RGBA(colors.LavandaEscuro))

	for i, title := range titles {
		headers[i] = canvas.NewText(title, theme.ForegroundColor())
		headers[i].TextStyle.Bold = true
		headers[i].Color = colors.LavandaClaro
		headers[i].Alignment = fyne.TextAlignCenter
	}

	productGrid := container.NewGridWithColumns(len(titles))
	for i := range headers {
		productGrid.Add(container.NewStack(background, headers[i]))

	}

	return productGrid

}

func createLabel(labelText string) *fyne.Container {
	background := canvas.NewRectangle(colors.LavandaClaro)

	Label := canvas.NewText(labelText, theme.ForegroundColor())
	Label.Color = colors.RoxoSuave
	Label.Alignment = fyne.TextAlignCenter

	return container.NewStack(background, Label)
}

func createLineToList(list *fyne.Container, product domain.Product) {
	list.Add(createLabel(product.Name))
	list.Add(createLabel(fmt.Sprintf("%.2f", product.Value)))
	list.Add(createLabel(fmt.Sprintf("%d", product.Quantity)))
	list.Add(createLabel(fmt.Sprintf("%.2f", product.TotalValue)))
	list.Add(createLabel(product.Market))
}

func createTotalLineToList(list *fyne.Container, products []domain.Product) {
	var values = make([]float64, len(products))
	for i, product := range products {
		values[i] = product.TotalValue
	}

	totalValue := utils.Sum(values)

	list.Add(createLabel("Total"))
	list.Add(createLabel(""))
	list.Add(createLabel(""))
	list.Add(createLabel(""))
	list.Add(createLabel(fmt.Sprintf("%.2f", totalValue)))
}

func fillListToShow(list *fyne.Container, products []domain.Product) {

	for _, product := range products {
		createLineToList(list, product)
	}

	createTotalLineToList(list, products)
}

func (g *GUIService) ListProductsByFilter(listContainer *fyne.Container, filter string) {
	listContainer.Objects = nil

	productGrid := createListToShow("Nome", "Valor", "Quantidade", "Total", "Mercado")

	products, err := g.UserService.GetProductsByFilter(filter)
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh()
		return
	}

	fillListToShow(productGrid, products)

	listContainer.Add(productGrid)
	listContainer.Refresh()
}

func (g *GUIService) ListMarkets(listContainer *fyne.Container, filter string) {
	listContainer.Objects = nil

	productGrid := createListToShow("Mercado")

	markets, err := g.UserService.GetMarkets()
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh()
		return
	}

	for _, market := range markets {
		productGrid.Add(createLabel(market))

	}

	listContainer.Add(productGrid)
	listContainer.Refresh()
}
