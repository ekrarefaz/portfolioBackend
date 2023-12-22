package services

import (
	"github.com/ekrarefaz/portfolioBackend/interfaces"
	"github.com/ekrarefaz/portfolioBackend/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	repo interfaces.IProjectRepository
	DB   *gorm.DB
}

// NewProjectService creates a new Project Service Object
func NewProjectService(repo interfaces.IProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// Create make a new project entry in the DB
func (s *ProjectService) Create(project *models.Project) error {
	return s.repo.Create(project)
}

// GetByID gets a project from the database given an id
func (s *ProjectService) GetByID(id uint) (*models.Project, error) {
	return s.repo.GetByID(id)
}

// GetAll gets all the existing projects
func (s *ProjectService) GetAll() ([]models.Project, error) {
	return s.repo.GetAll()
}

// DeleteByID deletes a project given an id
func (s *ProjectService) DeleteByID(id uint) error {
	return s.repo.DeleteByID(id)
}
