package repository

import "hexal_manager/pkg/domain/entities"

type ModuleRepository interface {
	GetById(id string) (*entities.Module, error)
	GetAll() ([]*entities.Module, error)
	AddOperations(operations []int64) (*entities.Module, error)
	Create(module *entities.Module) (*entities.Module, error)
	Update(module *entities.Module) (*entities.Module, error)
	Delete(id string) error
}
