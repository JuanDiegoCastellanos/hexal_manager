package test

import (
	"hexal_manager/app/domain/entities"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func createRandomModule(t *testing.T) *entities.Module {
	temporalModule := entities.NewModule(int64(fake.Day()), fake.UserName())
	require.NotEqualf(t, temporalModule.ID, int64(0), "module id should not be empty")
	require.NotEqual(t, temporalModule.Name, "")
	require.NotEmptyf(t, temporalModule, "The module can't be empty")
	return temporalModule
}

func TestCreateModule(t *testing.T) {
	createRandomModule(t)
}

func TestAddOperationModule(t *testing.T) {
	newOp := createRandomOperation(t)
	newMod := createRandomModule(t)
	arrayOp := newMod.AddOperation(*newOp)
	require.NotEmpty(t, arrayOp)
	require.True(t, len(arrayOp) > 0)
	require.NotEqual(t, "", arrayOp[0].Name)
}
