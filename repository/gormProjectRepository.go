// services/gormProjectRepository.go

package repository

import (
	"github.com/ekrarefaz/portfolioBackend/models"
	"gorm.io/gorm"
)

// GormProjectRepository is the GORM implementation of IProjectRepository
type GormProjectRepository struct {
	db *gorm.DB
}

// NewGormProjectRepository creates a new repository with a given GORM DB connection
func NewGormProjectRepository(db *gorm.DB) *GormProjectRepository {
	return &GormProjectRepository{db: db}
}

func (r *GormProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *GormProjectRepository) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, id).Error
	return &project, err
}

func (r *GormProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *GormProjectRepository) DeleteByID(id uint) error {
	err := r.db.Delete(&models.Project{}, id).Error
	return err
}
