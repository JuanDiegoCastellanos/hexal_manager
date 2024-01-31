package repository

import "hexal_manager/pkg/domain/entities"

type VehicleRepository interface {
	ListAll() (*[]entities.Vehicle, error)
	GetVehicle(id string) (*entities.Vehicle, error)
	Create(vh *entities.Vehicle) (*entities.Vehicle, error)
	Update(vh *entities.Vehicle) (*entities.Vehicle, error)
	Delete(id string) error
}
