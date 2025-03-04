package gui

import (
	"bronze/internal/application/services"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewUserList retorna o conteúdo da tela de listagem de usuários
func NewUserList(userService *services.UserService, w fyne.Window) fyne.CanvasObject {
	listContainer := container.NewVBox()
	refreshList := func() {
		RefreshUserList(userService, listContainer)
	}

	insertButton := widget.NewButton("Inserir Produto", func() {
		w.SetContent(NewCadastroForm(userService, w))
	})

	// Botão para voltar à tela principal
	reportButton := widget.NewButton("Gerar Relatorio", func() {
	})

	// Inicializa a lista com os usuários atuais
	refreshList()

	// Organiza os componentes na tela
	return container.NewVBox(
		listContainer,
		insertButton,
		reportButton,
	)
}

func RefreshUserList(userService *services.UserService, listContainer *fyne.Container) {
	// Limpa a lista antes de atualizar
	listContainer.Objects = nil

	// Obtém a lista de usuários
	users, err := userService.GetUsers()
	if err != nil {
		listContainer.Add(widget.NewLabel("Erro ao carregar usuários"))
		listContainer.Refresh() // Atualiza a exibição
		return
	}

	// Adiciona os usuários à lista
	for _, user := range users {
		listContainer.Add(widget.NewLabel(fmt.Sprintf("Produto: %s, Valor: %d", user.Name, user.Age)))
	}

	// Atualiza a exibição
	listContainer.Refresh()
}
