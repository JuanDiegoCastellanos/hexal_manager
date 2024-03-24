package usecases

import "hexal_manager/pkg/domain/entities"

type ModuleUseCase interface {
	NewModel() *entities.Module
	AddOperations() *entities.Module
	UpdateModule() *entities.Module
	DeleteModule() bool
	GetById() *entities.Module
	GetAll() []*entities.Module
}
type moduleUseCaseImpl struct {
}
