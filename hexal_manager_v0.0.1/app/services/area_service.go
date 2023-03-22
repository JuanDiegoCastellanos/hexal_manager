package services

import (
	"hexal_manager_v0.0.1/app/models"
	"hexal_manager_v0.0.1/app/repositories"
)

type AreaService struct {
	Repository *repositories.AreaRepository
}

func NewAreaService(areaRepo *repositories.AreaRepository) *AreaService {
	return &AreaService{Repository: areaRepo}
}
func (s *AreaService) CreateArea(area *models.Area) error {
	return s.Repository.Create(area)
}
func (s *AreaService) UpdateArea(area *models.Area) error {
	return s.Repository.Update(area)
}
func (s *AreaService) DeleteArea(area *models.Area) error {
	return s.Repository.Delete(area)
}
func (s *AreaService) FindAreaById(id uint) models.Area {
	return s.Repository.FindById(id)
}
func (s *AreaService) FindAllArea() ([]models.Area, error) {
	return s.Repository.FindAll()
}
