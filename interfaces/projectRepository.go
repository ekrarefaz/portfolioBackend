package interfaces

import "github.com/ekrarefaz/portfolioBackend/models"

// IProjectRepository defines the interface for project database operations
type IProjectRepository interface {
	Create(project *models.Project) error
	GetByID(id uint) (*models.Project, error)
	GetAll() ([]models.Project, error)
	DeleteByID(id uint) error
}
