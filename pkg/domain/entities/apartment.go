package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Apartment struct {
	ID              int64              `json:"id"`
	ApartmentNumber string             `json:"apartment_number"`
	OwnerId         int64              `json:"owner_id"`
	Size            int                `json:"size"`
	RoomCount       int                `json:"room_count"`
	BathRoomCount   int                `json:"bath_room_count"`
	DepositNumber   int                `json:"deposit_number"`
	State           helpers.EnumEntity `json:"state"`
	Cost            float64            `json:"cost"`
	RentalCost      float64            `json:"rental_cost"`
	CreatedAt       time.Time          `json:"created_at"`
}
