package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"hexal_manager/util/helpers"
	"testing"
)

func createNewEnum(t *testing.T) *helpers.EnumAdapter {
	stateEnum := helpers.NewEnumAdapter([]helpers.ValueEnum{})
	stateEnum.AddValue(*helpers.NewValueEnum(2, "busy"))
	stateEnum.AddValue(*helpers.NewValueEnum(3, "remodeling"))
	stateEnum.AddValue(*helpers.NewValueEnum(1, "free"))
	require.NotNil(t, stateEnum)
	require.NotEmpty(t, stateEnum.Values)
	require.Len(t, stateEnum.Values, 3)
	fmt.Println(stateEnum.Values)
	return stateEnum
}

func TestCreateEnum(t *testing.T) {
	createNewEnum(t)
}
