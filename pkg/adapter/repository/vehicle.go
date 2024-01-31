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

func (vhr *vehicleRepository) GetVehicle(id string) (*entities.Vehicle, error) {
	var vehicle = entities.Vehicle{}

	if err := vhr.db.First(&vehicle, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil

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

func (vhr *vehicleRepository) Update(vh *entities.Vehicle) (*entities.Vehicle, error) {
	var vehicle = entities.Vehicle{}

	err := vhr.db.First(&vehicle, "id = ?", vh.LicensePlate).Updates(&vehicle).Error

	if err != nil {
		return nil, err
	}
	return &vehicle, nil

}
func (vhr *vehicleRepository) Delete(id string) error {

	var vehicle = entities.Vehicle{}

	err := vhr.db.First(&vehicle, " id = ?", id).Error

	if err != nil {
		return err
	}
	err = vhr.db.Delete(&vehicle).Error
	if err != nil {
		return err
	}
	return nil

}
