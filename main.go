package main

import (
	"fmt"
	"github.com/ekrarefaz/portfolioBackend/controllers"
	"github.com/ekrarefaz/portfolioBackend/database"
	"github.com/ekrarefaz/portfolioBackend/repository"
	"github.com/ekrarefaz/portfolioBackend/routes"
	"github.com/ekrarefaz/portfolioBackend/services"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load Environment Variables using GoDotEnv
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading Environment Variables %v", err)
	}

	// Instantiate Database Connection
	db := database.ConnectDatabase()

	// Initialize the relevant repository
	projectRepo := repository.NewGormProjectRepository(db)
	aboutRepo := repository.NewGormAboutRepository(db)

	// Init services with the relevant repository
	projectService := services.NewProjectService(projectRepo)
	aboutService := services.NewAboutService(aboutRepo)

	// Init controllers with the relevant service
	projectController := controllers.NewProjectController(projectService)
	aboutController := controllers.NewAboutController(aboutService)

	// Init Router with the controller
	router := routes.RouterInit(projectController, aboutController)

	// Start the server
	err = router.Run(":8888")
	if err != nil {
		panic("Failed to run server: " + err.Error())
	}
	fmt.Println("Server running on port 8888")
}
