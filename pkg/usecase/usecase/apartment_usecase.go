package usecases

import (
	"errors"
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/pkg/usecase/repository"
)

type ApartmentUseCase interface {
	GetById(id string) (*entities.Apartment, error)
	GetAll() ([]*entities.Apartment, error)
	CreateApartment(apartment *entities.Apartment) (*entities.Apartment, error)
	UpdateApartment(apartment *entities.Apartment) (*entities.Apartment, error)
	DeleteApartment(id string) error
}
type apartmentUseCaseImpl struct {
	apartmentRepository repository.ApartmentRepository
	dBRepository        repository.DBRepository
}

func (auc *apartmentUseCaseImpl) GetById(id string) (*entities.Apartment, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	apa, err := auc.apartmentRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return apa, nil
}

func (auc *apartmentUseCaseImpl) GetAll() ([]*entities.Apartment, error) {
	listApa, err := auc.apartmentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return listApa, nil
}
func (auc *apartmentUseCaseImpl) CreateApartment(apartment *entities.Apartment) (*entities.Apartment, error) {
	if apartment == nil {
		return nil, errors.New("apartment cannot be null")
	}
	apa, err := auc.apartmentRepository.Create(apartment)
	if err != nil {
		return nil, err
	}
	return apa, nil
}
func (auc *apartmentUseCaseImpl) UpdateApartment(apartment *entities.Apartment) (*entities.Apartment, error) {
	if apartment == nil {
		return nil, errors.New("apartment cannot be null")
	}
	apa, err := auc.apartmentRepository.Update(apartment)
	if err != nil {
		return nil, err
	}
	return apa, nil
}
func (auc *apartmentUseCaseImpl) DeleteApartment(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	err := auc.apartmentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
