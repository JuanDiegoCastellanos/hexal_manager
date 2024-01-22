package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Service struct {
	ID           int64
	Name         string
	TotalCost    float64
	GeneratedAt  time.Time
	ExpiringDate time.Time
	Type         helpers.EnumEntity
	OwnerID      int64
	ApartmentID  int64
}

func NewService(id int64, name string, totalCost float64,
	generatedAt, expiringDate time.Time, typeService helpers.EnumEntity,
	ownerID, apartmentID int64) *Service {

	return &Service{
		ID:           id,
		Name:         name,
		TotalCost:    totalCost,
		GeneratedAt:  generatedAt,
		ExpiringDate: expiringDate,
		Type:         typeService,
		OwnerID:      ownerID,
		ApartmentID:  apartmentID,
	}

}
