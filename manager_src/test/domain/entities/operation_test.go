package test

import (
	"manager_src/app/domain/entities"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func createRandomOperation(t *testing.T) *entities.Operation {
	randomOperation := entities.NewOperation(int64(fake.Day()), fake.ProductName(), make([]entities.Rol, 0))
	require.NotZero(t, randomOperation.ID)
	require.NotEmpty(t, randomOperation.Name)
	require.NotEqual(t, "", randomOperation.Name)
	return randomOperation
}

func TestNewOperation(t *testing.T) {
	createRandomOperation(t)
}
