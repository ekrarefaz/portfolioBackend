package controllers

import (
	"net/http"
	"strconv"

	"github.com/ekrarefaz/portfolioBackend/models"
	"github.com/ekrarefaz/portfolioBackend/services"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service *services.ProjectService
}

// NewProjectController creates a new controller with given service
func NewProjectController(s *services.ProjectService) *ProjectController {
	return &ProjectController{
		service: s,
	}
}

// CreateProjectHandler handles the creation of a new project
func (pc *ProjectController) CreateProjectHandler(c *gin.Context) {
	var project models.Project

	err := c.ShouldBindJSON(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = pc.service.Create(&project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// GetProjectByIDHandler handles the retrieval of a project by ID
func (pc *ProjectController) GetProjectByIDHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	project, err := pc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetAllProjectsHandler handles the retrieval of all projects
func (pc *ProjectController) GetAllProjectsHandler(c *gin.Context) {
	projects, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// DeleteProjectByIDHandler handles the deletion of a project by ID
func (pc *ProjectController) DeleteProjectByIDHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = pc.service.DeleteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
