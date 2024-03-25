package usecases

import (
	"errors"
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/pkg/usecase/repository"
)

type ModuleUseCase interface {
	GetById(id string) (*entities.Module, error)
	GetAll() ([]*entities.Module, error)
	AddOperations(operations []int64) (*entities.Module, error)
	CreateModule(module *entities.Module) (*entities.Module, error)
	UpdateModule(module *entities.Module) (*entities.Module, error)
	DeleteModule(id string) error
}

type moduleUseCaseImpl struct {
	moduleRepository repository.ModuleRepository
	dbRepository     repository.DBRepository
}

func (muc *moduleUseCaseImpl) GetById(id string) (*entities.Module, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	mod, err := muc.moduleRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return mod, nil
}
func (muc *moduleUseCaseImpl) GetAll() ([]*entities.Module, error) {
	modsList, err := muc.moduleRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return modsList, err
}
func (muc *moduleUseCaseImpl) AddOperations(operations []int64) (*entities.Module, error) {
	if len(operations) > 0 {
		return nil, errors.New("the list cannot be empty")
	}
	mod, err := muc.moduleRepository.AddOperations(operations)
	if err != nil {
		return nil, err
	}
	return mod, nil
}
func (muc *moduleUseCaseImpl) CreateModule(module *entities.Module) (*entities.Module, error) {
	if module == nil {
		return nil, errors.New("module cannot be null or empty")
	}
	mod, err := muc.moduleRepository.Create(module)
	if err != nil {
		return nil, err
	}
	return mod, nil
}
func (muc *moduleUseCaseImpl) UpdateModule(module *entities.Module) (*entities.Module, error) {
	if module == nil {
		return nil, errors.New("module cannot be bull or empty")
	}
	mod, err := muc.moduleRepository.Update(module)
	if err != nil {
		return nil, err
	}
	return mod, err
}

func (muc *moduleUseCaseImpl) DeleteModule(id string) error {
	if id == "" {
		return errors.New("")
	}
	err := muc.moduleRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
