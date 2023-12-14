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
	database.ConnectDatabase()

	// Init project services and controllers
	projectService := services.NewProjectService(database.DB)
	projectController := controllers.NewProjectController(projectService)

	// Init Router
	router := routes.RouterInit(projectController)

	// Start the server
	err = router.Run(":8888")
	if err != nil {
		panic("Failed to run server")
	}
	fmt.Println("Server running...")
}
