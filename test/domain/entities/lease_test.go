package test

import (
	"github.com/stretchr/testify/require"
	"hexal_manager/app/domain/entities"
	"hexal_manager/util/helpers"
	"testing"
	"time"
)

func newLease(t *testing.T) *entities.Lease {
	leaseTest := entities.NewLease(
		helpers.RandomInt(1, 20), time.Now(),
		time.Now(), float64(helpers.RandomMoney()), helpers.RandomInt(1, 20))
	require.NotEmpty(t, leaseTest.ID)
	require.NotEmpty(t, leaseTest.StartDate)
	require.NotEmpty(t, leaseTest.EndDate)
	require.NotEmpty(t, leaseTest.PaymentMethodID)
	require.NotEmpty(t, leaseTest.TotalCost)

	return leaseTest
}

func TestNewLease(t *testing.T) {
	newLease(t)
}
