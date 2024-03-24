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

func (vuc *vehicleUseCase) GetVehicle(id string) (*entities.Vehicle, error) {
	vehicle, err := vuc.vehicleRepository.GetVehicle(id)

	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (vuc *vehicleUseCase) GetAll(id string) ([]*entities.Vehicle, error) {
	vehicles, err := vuc.vehicleRepository.ListAll()

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (vuc *vehicleUseCase) Create(vh *entities.Vehicle) (*entities.Vehicle, error) {

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
func (vuc *vehicleUseCase) Update(vh *entities.Vehicle) (*entities.Vehicle, error) {
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

func (vuc *vehicleUseCase) Delete(id string) error {
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
