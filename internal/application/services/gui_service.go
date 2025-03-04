package services

import (
	"fmt"

	"fyne.io/fyne/v2"
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
	// Limpa a lista antes de atualizar
	listContainer.Objects = nil

	// Obtém a lista de usuários
	products, err := g.UserService.GetProducts()
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar Produtos"))
		listContainer.Refresh() // Atualiza a exibição
		return
	}

	// Adiciona os usuários à lista
	for _, product := range products {
		listContainer.Add(widget.NewLabel(fmt.Sprintf("Produto: %s Valor: %.2f", product.Name, product.Value)))
	}

	// Atualiza a exibição
	listContainer.Refresh()
}
