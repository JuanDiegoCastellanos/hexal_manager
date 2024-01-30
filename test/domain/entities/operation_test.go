package test

import (
	"hexal_manager/pkg/domain/entities"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func createRandomOperation(t *testing.T) *entities.Operation {
	randomOperation := entities.NewOperation(int64(fake.Day()), fake.ProductName(), make([]entities.Rol, 1))
	require.NotZero(t, randomOperation.ID)
	require.NotEmpty(t, randomOperation.Name)
	require.NotEqual(t, "", randomOperation.Name)
	require.NotEmpty(t, randomOperation.Roles)

	return randomOperation
}

func TestNewOperation(t *testing.T) {
	op := createRandomOperation(t)
	require.NotNil(t, op)
}

func TestAddRol(t *testing.T) {

}
