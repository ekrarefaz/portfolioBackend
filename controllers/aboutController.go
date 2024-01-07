package controllers

import (
	"github.com/ekrarefaz/portfolioBackend/models"
	"github.com/ekrarefaz/portfolioBackend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AboutController struct {
	service *services.AboutService
}

func NewAboutController(s *services.AboutService) *AboutController {
	return &AboutController{service: s}
}

func (ac *AboutController) CreateAboutHandler(c *gin.Context) {
	var about models.About

	err := c.ShouldBindJSON(&about)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ac.service.Create(&about)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, about)
}

func (ac *AboutController) DeleteAboutHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = ac.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "About deleted"})

}

func (ac *AboutController) GetAboutHandler(c *gin.Context) {
	about, err := ac.service.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, about)

}
