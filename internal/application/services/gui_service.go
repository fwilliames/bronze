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

func createListToShow() *fyne.Container {

	background := canvas.NewRectangle(color.RGBA(colors.LavandaEscuro))

	header1 := canvas.NewText("Nome", theme.ForegroundColor())
	header1.TextStyle.Bold = true
	header1.Color = colors.LavandaClaro

	header2 := canvas.NewText("Preço", theme.ForegroundColor())
	header2.TextStyle.Bold = true
	header2.Color = colors.LavandaClaro

	header3 := canvas.NewText("Mes/Ano", theme.ForegroundColor())
	header3.TextStyle.Bold = true
	header3.Color = colors.LavandaClaro

	productGrid := container.NewGridWithColumns(3)
	productGrid.Add(container.NewStack(background, header1))
	productGrid.Add(container.NewStack(background, header2))
	productGrid.Add(container.NewStack(background, header3))

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

	productGrid := createListToShow()

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

	}

	listContainer.Add(productGrid)
	listContainer.Refresh()
}
