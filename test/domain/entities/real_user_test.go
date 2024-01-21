package test

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"hexal_manager/app/domain/entities"
	"hexal_manager/util/helpers"
	"testing"
	"time"
)

func newRealUser(t *testing.T) *entities.RealUser {
	realUserTest := entities.NewRealUser(
		helpers.RandomInt(1, 20),
		fake.UserName(),
		fake.LastName(),
		fake.DigitsN(12),
		fake.DigitsN(10),
		fake.JobTitle(),
		int(helpers.RandomInt(1, 10)),
		helpers.RandomInt(1, 20),
		helpers.RandomInt(1, 20),
		helpers.RandomInt(1, 20),
		helpers.RandomInt(1, 20),
		time.Now(),
	)

	require.NotEmpty(t, realUserTest.ID)
	require.NotEmpty(t, realUserTest.Name)
	require.NotEmpty(t, realUserTest.LastName)
	require.NotEmpty(t, realUserTest.Document)
	require.NotEmpty(t, realUserTest.PhoneNumber)
	require.NotEmpty(t, realUserTest.Occupation)
	require.NotEmpty(t, realUserTest.PetsCount)
	require.NotEmpty(t, realUserTest.ApartmentId)
	require.NotEmpty(t, realUserTest.VehicleId)
	require.NotEmpty(t, realUserTest.LeaseId)
	require.NotEmpty(t, realUserTest.UserId)
	require.NotEmpty(t, realUserTest.CreatedAt)

	return realUserTest
}

func TestCreateRealUser(t *testing.T) {
	newRealUser(t)
}
