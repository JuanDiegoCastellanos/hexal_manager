package entities

import "hexal_manager/util/helpers"

type Apartment struct {
	ID              int64
	ApartmentNumber string
	OwnerId         int64
	Size            int
	RoomCount       int
	BathRoomCount   int
	DepositNumber   int
	State           helpers.EnumEntity
}
