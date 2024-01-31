package controller

import (
	"hexal_manager/pkg/domain/entities"
	usecases "hexal_manager/pkg/usecase/usecase"
	"net/http"
)

type vehicleController struct {
	vehicleUseCase usecases.VehicleUseCase
}

type Vehicle interface {
	GetVehicles(ctx Context) error
	AddVehicle(ctx Context) error
}

func NewVehicleController(us usecases.VehicleUseCase) Vehicle {
	return &vehicleController{us}
}

func (vhc *vehicleController) GetVehicles(ctx Context) error {

	data, err := vhc.vehicleUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)

}
func (vhc *vehicleController) AddVehicle(ctx Context) error {
	var params = entities.Vehicle{}

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	vh, err := vhc.vehicleUseCase.Create(&params)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, vh)

}
