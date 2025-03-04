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
		listContainer.Objects = nil // Limpa a lista antes de atualizar
		users, err := userService.GetUsers()
		if err != nil {
			listContainer.Add(widget.NewLabel("Erro ao carregar usuários"))
			return
		}

		for _, user := range users {
			listContainer.Add(widget.NewLabel(fmt.Sprintf("Nome: %s, Idade: %d", user.Name, user.Age)))
		}

		listContainer.Refresh() // Atualiza a exibição
	}

	// Botão para recarregar os usuários
	refreshButton := widget.NewButton("Recarregar", func() {
		refreshList()
	})

	cadastroButton := widget.NewButton("Tela de Cadastro", func() {
		w.SetContent(NewCadastroForm(userService, w))
	})

	// Botão para voltar à tela principal
	backButton := widget.NewButton("Voltar", func() {
		w.SetContent(NewMainWindow(userService, w))
	})

	// Inicializa a lista com os usuários atuais
	refreshList()

	// Organiza os componentes na tela
	return container.NewVBox(
		widget.NewLabel("Lista de Usuários"),
		refreshButton,
		listContainer,
		cadastroButton,
		backButton,
	)
}
