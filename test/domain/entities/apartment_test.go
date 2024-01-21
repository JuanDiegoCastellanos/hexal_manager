package test

import (
	"github.com/icrowley/fake"
	"hexal_manager/app/domain/entities"
	"hexal_manager/util/helpers"
	"math/rand"
	"testing"
	"time"
)

func newApartment(t *testing.T) *entities.Apartment {
	apartmentTest := entities.NewApartment(
		helpers.RandomInt(1, 20),
		helpers.RandomInt(1, 20),
		fake.DigitsN(10),
		rand.Intn(1),
		rand.Intn(1),
		rand.Intn(1),
		rand.Intn(1),
		helpers.NewEnumAdapter([]helpers.ValueEnum{}),
		float64(helpers.RandomMoney()),
		float64(helpers.RandomMoney()),
		time.Now(),
	)
	return apartmentTest
}

func TestCreateApartment(t *testing.T) {
	newApartment(t)
}
