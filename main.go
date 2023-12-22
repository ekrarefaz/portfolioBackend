package main

import (
	"fmt"
	"github.com/ekrarefaz/portfolioBackend/controllers"
	"github.com/ekrarefaz/portfolioBackend/database"
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

	// Initialize the GORM repository
	gormRepo := &services.ProjectService{DB: db}

	// Init project services with the GORM repository
	projectService := services.NewProjectService(gormRepo)

	// Init project controller with the service
	projectController := controllers.NewProjectController(projectService)

	// Init Router with the controller
	router := routes.RouterInit(projectController)

	// Start the server
	err = router.Run(":8888")
	if err != nil {
		panic("Failed to run server: " + err.Error())
	}
	fmt.Println("Server running on port 8888")
}
