package services

import (
	"github.com/ekrarefaz/portfolioBackend/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	err := s.db.Find(&projects).Error
	return projects, err
}
