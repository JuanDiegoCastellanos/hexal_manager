package repository

import "hexal_manager/pkg/domain/entities"

type VehicleRepository interface {
	GetAll() ([]*entities.Vehicle, error)
	GetById(id string) (*entities.Vehicle, error)
	Create(vh *entities.Vehicle) (*entities.Vehicle, error)
	Update(vh *entities.Vehicle) (*entities.Vehicle, error)
	Delete(id string) error
}
