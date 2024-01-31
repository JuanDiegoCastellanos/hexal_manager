package usecases

import (
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/pkg/usecase/repository"
)

type vehicleUseCase struct {
	vehicleRepository repository.VehicleRepository
	dBRepository      repository.DBRepository
}

type VehicleUseCase interface {
	GetVehicle(id string) (*entities.Vehicle, error)
	GetAll() (*[]entities.Vehicle, error)
	Create(vh *entities.Vehicle) error
	Update(vh *entities.Vehicle) error
	Delete(id string) error
}

func (vhU *vehicleUseCase) GetVehicle(id string) (*entities.Vehicle, error) {
	vh, err := vhU.vehicleRepository.GetVehicle(id)
	if err != nil {
		return nil, err
	}

	return vh, nil
}
