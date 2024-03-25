package usecases

import (
	"errors"
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/pkg/usecase/repository"
)

type VehicleUseCase interface {
	GetById(id string) (*entities.Vehicle, error)
	GetAll() (*[]entities.Vehicle, error)
	CreateVehicle(vh *entities.Vehicle) (*entities.Vehicle, error)
	UpdateVehicle(vh *entities.Vehicle) (*entities.Vehicle, error)
	DeleteVehicle(id string) error
}
type vehicleUseCaseImpl struct {
	vehicleRepository repository.VehicleRepository
	dBRepository      repository.DBRepository
}

func (vuc *vehicleUseCaseImpl) GetById(id string) (*entities.Vehicle, error) {
	vehicle, err := vuc.vehicleRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (vuc *vehicleUseCaseImpl) GetAll(id string) ([]*entities.Vehicle, error) {
	vehicles, err := vuc.vehicleRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (vuc *vehicleUseCaseImpl) CreateVehicle(vh *entities.Vehicle) (*entities.Vehicle, error) {
	data, err := vuc.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		vehicle, err := vuc.vehicleRepository.Create(vh)
		if err != nil {
			return nil, err
		}
		return vehicle, nil
	})
	vehicle, ok := data.(*entities.Vehicle)
	if !ok {
		return nil, errors.New("cast error")
	}
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
func (vuc *vehicleUseCaseImpl) UpdateVehicle(vh *entities.Vehicle) (*entities.Vehicle, error) {
	data, err := vuc.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		vh, err := vuc.vehicleRepository.Update(vh)
		if err != nil {
			return nil, err
		}
		return vh, nil
	})
	updatedVehicle, ok := data.(*entities.Vehicle)
	if !ok {
		return nil, errors.New("cast error")
	}
	if err != nil {
		return nil, err
	}
	return updatedVehicle, nil
}

func (vuc *vehicleUseCaseImpl) DeleteVehicle(id string) error {
	_, err := vuc.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		err := vuc.vehicleRepository.Delete(id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
