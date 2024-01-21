package test

import (
	"github.com/icrowley/fake"
	"hexal_manager/app/domain/entities"
	"hexal_manager/util/helpers"
	"testing"
)

func newVehicle(t *testing.T) *entities.Vehicle {
	vehicleTest := entities.NewVehicle(
		helpers.RandomInt(1, 20),
		fake.CharactersN(10),
		fake.CharactersN(10),
		helpers.NewEnumAdapter([]helpers.ValueEnum{}),
		helpers.RandomInt(1, 20),
	)
	return vehicleTest
}
