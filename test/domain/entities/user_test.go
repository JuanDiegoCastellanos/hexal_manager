package test

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"hexal_manager/app/domain/entities"
	"hexal_manager/util/helpers"
	"testing"
)

func newUser(t *testing.T) *entities.User {
	userTest := entities.NewUser(helpers.RandomInt(1, 20),
		fake.UserName(), fake.EmailAddress(),
		fake.Password(8, 10, true, true, true),
		helpers.RandomInt(1, 20))

	require.NotEmpty(t, userTest)
	require.Greater(t, userTest.ID, int64(0))
	require.NotEmpty(t, userTest.NickName)
	require.NotEmpty(t, userTest.Email)
	require.NotEmpty(t, userTest.Password)
	require.Greater(t, userTest.RoleId, int64(0))

	return userTest
}

func TestCreateUser(t *testing.T) {
	newUser(t)
}
