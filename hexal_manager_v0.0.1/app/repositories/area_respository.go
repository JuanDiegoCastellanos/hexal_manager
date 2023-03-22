package repositories

import (
	"gorm.io/gorm"
	"hexal_manager_v0.0.1/app/models"
)

type AreaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *AreaRepository {
	return &AreaRepository{db: db}
}

func (ar *AreaRepository) Create(area *models.Area) error {

	return ar.db.Create(area).Error
}

func (ar *AreaRepository) Update(area *models.Area) error {
	return ar.db.Save(area).Error
}

func (ar *AreaRepository) Delete(area *models.Area) error {
	return ar.db.Delete(area).Error
}
func (ar *AreaRepository) FindById(id uint) models.Area {
	var area models.Area
	err := ar.db.First(&area, id).Error
	if err != nil {
		return models.Area{}
	}
	return area
}
func (ar *AreaRepository) FindAll() ([]models.Area, error) {
	var list_areas []models.Area
	err := ar.db.Find(&list_areas).Error
	if err != nil {
		return nil, err
	}
	return list_areas, nil

}
