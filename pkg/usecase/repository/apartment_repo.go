package repository

import "hexal_manager/pkg/domain/entities"

type ApartmentRepository interface {
	GetById(id string) (*entities.Apartment, error)
	GetAll() ([]*entities.Apartment, error)
	Create(apartment *entities.Apartment) (*entities.Apartment, error)
	Update(apartment *entities.Apartment) (*entities.Apartment, error)
	Delete(id string) error
}
