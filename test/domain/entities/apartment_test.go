package test

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
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
		rand.Int(),
		rand.Int(),
		rand.Int(),
		rand.Int(),
		helpers.NewEnumAdapter([]helpers.ValueEnum{}),
		float64(helpers.RandomMoney()),
		float64(helpers.RandomMoney()),
		time.Now(),
	)
	require.NotEmpty(t, apartmentTest.ID)
	require.NotEmpty(t, apartmentTest.ApartmentNumber)
	require.NotEmpty(t, apartmentTest.OwnerId)
	require.NotEmpty(t, apartmentTest.Cost)
	require.NotEmpty(t, apartmentTest.DepositNumber)
	require.NotEmpty(t, apartmentTest.RoomCount)
	require.NotEmpty(t, apartmentTest.Size)
	require.NotEmpty(t, apartmentTest.State)
	require.NotEmpty(t, apartmentTest.RentalCost)
	require.NotEmpty(t, apartmentTest.BathRoomCount)
	require.NotEmpty(t, apartmentTest.CreatedAt)

	return apartmentTest
}

func TestCreateApartment(t *testing.T) {
	newApartment(t)
}
