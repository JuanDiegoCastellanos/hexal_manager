package test

import (
	"manager_src/app/domain/entities"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func createNewRol(t *testing.T) *entities.Rol {
	rol := entities.NewRol(int64(fake.Day()), fake.UserName(), make([]entities.Operation, 1))
	require.NotEmpty(t, rol)
	require.NotEqual(t, rol.ID, 0)
	require.NotEmpty(t, rol.ID)
	require.NotEmpty(t, rol.Name)
	require.NotEmpty(t, rol.Operations)
	require.NotNil(t, rol.Operations)
	return rol
}

func TestCreateNewRol(t *testing.T) {
	rol := createNewRol(t)
	require.NotNil(t, rol)
}
