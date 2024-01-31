package usecases

import (
	"errors"
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
	Create(vh *entities.Vehicle) (*entities.Vehicle, error)
	Update(vh *entities.Vehicle) (*entities.Vehicle, error)
	Delete(id string) error
}

func (vhu *vehicleUseCase) GetVehicle(id string) (*entities.Vehicle, error) {
	vehicle, err := vhu.vehicleRepository.GetVehicle(id)

	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (vhU *vehicleUseCase) GetAll(id string) ([]*entities.Vehicle, error) {
	vehicles, err := vhU.vehicleRepository.ListAll()

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (vhu *vehicleUseCase) Create(vh *entities.Vehicle) (*entities.Vehicle, error) {

	data, err := vhu.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		vehicle, err := vhu.vehicleRepository.Create(vh)
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
func (vhu *vehicleUseCase) Update(vh *entities.Vehicle) (*entities.Vehicle, error) {
	data, err := vhu.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		vh, err := vhu.vehicleRepository.Update(vh)
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

func (vhu *vehicleUseCase) Delete(id string) error {
	_, err := vhu.dBRepository.TransactionTX(func(i interface{}) (interface{}, error) {
		err := vhu.vehicleRepository.Delete(id)
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
