package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type ParkingSpace struct {
	ParkingNumber    int64              `json:"parking_number"`
	ParkingLot       int64              `json:"parking_lot"`
	ApartmentID      int64              `json:"apartments_id"`
	ParkingSpaceType helpers.EnumEntity `json:"parking_space_type"`
	State            helpers.EnumEntity `json:"state"`
	CheckIn          time.Time          `json:"check_in"`
	CheckOut         time.Time          `json:"check_out"`
	VehicleID        int64              `json:"vehicle_id"`
}

func NewParkingSpace(parkingNumber, parkingLot, apartmentID, vehicleID int64,
	parkingSpaceType, state helpers.EnumEntity, checkIn, checkOut time.Time) *ParkingSpace {

	return &ParkingSpace{
		ParkingNumber:    parkingNumber,
		ParkingLot:       parkingLot,
		ApartmentID:      apartmentID,
		ParkingSpaceType: parkingSpaceType,
		State:            state,
		CheckIn:          checkIn,
		CheckOut:         checkOut,
		VehicleID:        vehicleID,
	}
}
