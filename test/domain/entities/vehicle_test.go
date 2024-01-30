package test

import (
	"hexal_manager/pkg/domain/entities"
	"hexal_manager/util/helpers"
	"testing"

	"github.com/icrowley/fake"
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

func TestCreateVehicle(t *testing.T) {
	newVehicle(t)
}
