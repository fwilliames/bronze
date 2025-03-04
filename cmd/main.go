package main

import (
	"bronze/internal/adapters/database"
	"bronze/internal/adapters/gui"
	"bronze/internal/application/services"
	"log"
)

func main() {
	// Inicializa o repositório (banco de dados)
	repo, err := database.NewSQLiteRepository()
	if err != nil {
		log.Println("Erro ao iniciar banco:", err) // Apenas um log, sem travar
		return
	}

	// Inicializa o serviço de usuários
	userService := services.NewUserService(repo)

	// Inicializa a interface gráfica
	gui.StartApp(userService)
}
