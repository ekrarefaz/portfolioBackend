package controllers

import (
	"github.com/ekrarefaz/portfolioBackend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectController struct {
	projectService *services.ProjectService
}

func NewProjectController(projectService *services.ProjectService) *ProjectController {
	return &ProjectController{projectService: projectService}
}

func (pc *ProjectController) GetProjectsHandler(c *gin.Context) {
	projects, err := pc.projectService.GetAllProjects()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}
	c.JSON(http.StatusOK, projects)
}
