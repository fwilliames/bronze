package main

import (
	"bronze/internal/adapters/database"
	"bronze/internal/adapters/gui"
	"bronze/internal/application/services"
	"log"
)

func main() {

	repo, err := database.NewSQLiteRepository()
	if err != nil {
		log.Println("Erro ao iniciar banco:", err)
		return
	}

	userService := services.NewUserService(repo)
	guiService := services.NewGUIService(userService)

	gui.StartApp(guiService)
}
