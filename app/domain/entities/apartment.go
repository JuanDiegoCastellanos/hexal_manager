package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Apartment struct {
	ID              int64              `json:"id,omitempty"`
	ApartmentNumber string             `json:"apartment_number,omitempty"`
	OwnerId         int64              `json:"owner_id,omitempty"`
	Size            int                `json:"size,omitempty"`
	RoomCount       int                `json:"room_count,omitempty"`
	BathRoomCount   int                `json:"bath_room_count,omitempty"`
	DepositNumber   int                `json:"deposit_number,omitempty"`
	State           helpers.EnumEntity `json:"state,omitempty"`
	Cost            float64            `json:"cost,omitempty"`
	RentalCost      float64            `json:"rental_cost,omitempty"`
	CreatedAt       time.Time          `json:"created_at"`
}

func NewApartment(id, ownerId int64, apartmentNumber string,
	size, roomCount, bathRoomCount, depositNumber int,
	stateEnum helpers.EnumEntity, cost, rentalCost float64,
	createdAt time.Time) *Apartment {
	return &Apartment{
		ID:              id,
		ApartmentNumber: apartmentNumber,
		OwnerId:         ownerId,
		Size:            size,
		RoomCount:       roomCount,
		BathRoomCount:   bathRoomCount,
		DepositNumber:   depositNumber,
		State:           stateEnum,
		Cost:            cost,
		RentalCost:      rentalCost,
		CreatedAt:       createdAt,
	}
}
