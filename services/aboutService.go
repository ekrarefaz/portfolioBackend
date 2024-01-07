package services

import (
	"github.com/ekrarefaz/portfolioBackend/interfaces"
	"github.com/ekrarefaz/portfolioBackend/models"
)

type AboutService struct {
	repo interfaces.IAboutRepository
}

func NewAboutService(repo interfaces.IAboutRepository) *AboutService {
	return &AboutService{repo: repo}
}

func (s *AboutService) Create(about *models.About) error {
	return s.repo.Create(about)
}

func (s *AboutService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AboutService) Update(about *models.About) error {
	return s.repo.Update(about)
}

func (s *AboutService) Get() (*models.About, error) {
	return s.repo.Get()
}
