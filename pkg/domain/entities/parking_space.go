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
