package services

import (
	colors "bronze/internal/config/colors"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GUIService fornece funcionalidades para manipulação da interface gráfica
type GUIService struct {
	UserService *UserService
	Filters     Filters
}

type Filters struct {
	Data   string
	Market string
}

// NewGUIService cria uma instância do serviço de GUI
func NewGUIService(userService *UserService) *GUIService {
	return &GUIService{UserService: userService}
}

// RefreshUserList atualiza a lista de usuários na interface gráfica
func (g *GUIService) RefreshUserList(listContainer *fyne.Container) {

	listContainer.Objects = nil

	productGrid := createListToShow()

	products, err := g.UserService.GetProducts()
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh() // Atualiza a exibição
		return
	}

	for _, product := range products {
		productGrid.Add(createLabel(product.Name))
		productGrid.Add(createLabel(fmt.Sprintf("%.2f", product.Value)))
		productGrid.Add(createLabel(product.Data))

	}

	listContainer.Add(productGrid)
	listContainer.Refresh()
}

func createListToShow(titles ...string) *fyne.Container {

	headers := make([]*canvas.Text, len(titles))
	background := canvas.NewRectangle(color.RGBA(colors.LavandaEscuro))

	for i, title := range titles {
		headers[i] = canvas.NewText(title, theme.ForegroundColor())
		headers[i].TextStyle.Bold = true
		headers[i].Color = colors.LavandaClaro
	}

	productGrid := container.NewGridWithColumns(len(titles))
	for i, _ := range headers {
		productGrid.Add(container.NewStack(background, headers[i]))

	}

	return productGrid

}

func createLabel(labelText string) *fyne.Container {
	background := canvas.NewRectangle(colors.LavandaClaro)

	Label := canvas.NewText(labelText, theme.ForegroundColor())
	Label.Color = colors.RoxoSuave

	return container.NewStack(background, Label)
}

func (g *GUIService) ListProductsByFilter(listContainer *fyne.Container, filter string) {
	listContainer.Objects = nil

	productGrid := createListToShow("Nome", "Valor", "Mes/Ano", "Mercado")

	products, err := g.UserService.GetProductsByFilter(filter)
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh() // Atualiza a exibição
		return
	}

	for _, product := range products {
		productGrid.Add(createLabel(product.Name))
		productGrid.Add(createLabel(fmt.Sprintf("%.2f", product.Value)))
		productGrid.Add(createLabel(product.Data))
		productGrid.Add(createLabel(product.Market))

	}

	listContainer.Add(productGrid)
	listContainer.Refresh()
}

func (g *GUIService) ListMarkets(listContainer *fyne.Container, filter string) {
	listContainer.Objects = nil

	productGrid := createListToShow("Mercado")

	markets, err := g.UserService.GetMarkets()
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh() // Atualiza a exibição
		return
	}

	for _, market := range markets {
		productGrid.Add(createLabel(market))

	}

	listContainer.Add(productGrid)
	listContainer.Refresh()
}
