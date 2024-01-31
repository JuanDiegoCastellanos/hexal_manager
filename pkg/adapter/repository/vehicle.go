package repository

import (
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/pkg/usecase/repository"

	"gorm.io/gorm"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) repository.VehicleRepository {
	return &vehicleRepository{db}
}

func (vhr *vehicleRepository) ListAll() ([]*entities.Vehicle, error) {
	var vehicles = []*entities.Vehicle{}

	err := vhr.db.Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil

}
func (vhr *vehicleRepository) Create(vh *entities.Vehicle) (*entities.Vehicle, error) {
	if err := vhr.db.Create(vh).Error; err != nil {
		return nil, err
	}
	return vh, nil

}
