package interfaces

import "github.com/ekrarefaz/portfolioBackend/models"

type IAboutRepository interface {
	Create(about *models.About) error
	Delete(id uint) error
	Update(about *models.About) error
	Get() (*models.About, error)
}
