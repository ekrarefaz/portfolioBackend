package routes

import (
	"github.com/ekrarefaz/portfolioBackend/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterInit(
	projectController *controllers.ProjectController,
) *gin.Engine {

	// Use default gin router
	router := gin.Default()

	// Health Check Route
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Project Page Routes
	router.GET("/projects", projectController.GetProjectsHandler)

	return router
}
