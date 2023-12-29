package routes

import (
	"github.com/ekrarefaz/portfolioBackend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RouterInit(
	projectController *controllers.ProjectController,
) *gin.Engine {

	// Use default gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Or use AllowOrigins and list specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health Check Route
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Project Page Routes
	router.GET("/projects", projectController.GetAllProjectsHandler)
	router.POST("/projects", projectController.CreateProjectHandler)
	router.GET("/projects/:id", projectController.GetProjectByIDHandler)
	router.DELETE("/projects/:id", projectController.DeleteProjectByIDHandler) // Corrected to DELETE

	return router
}
