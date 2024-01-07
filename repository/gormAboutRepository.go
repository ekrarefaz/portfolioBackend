package repository

import (
	"github.com/ekrarefaz/portfolioBackend/models"
	"gorm.io/gorm"
)

type GormAboutRepository struct {
	db *gorm.DB
}

func NewGormAboutRepository(db *gorm.DB) *GormAboutRepository {
	return &GormAboutRepository{db: db}
}

func (r *GormAboutRepository) Create(about *models.About) error {
	return r.db.Create(about).Error
}

func (r *GormAboutRepository) Delete(id uint) error {
	err := r.db.Delete(&models.About{}, id).Error
	return err
}

func (r *GormAboutRepository) Update(about *models.About) error {
	err := r.db.Save(about).Error
	return err
}

func (r *GormAboutRepository) Get() (*models.About, error) {
	var about models.About
	err := r.db.First(&about).Error
	return &about, err
}
