package entities

import "hexal_manager/util/helpers"

type Vehicle struct {
	LicensePlate int64              `json:"license_plate"`
	Brand        string             `json:"brand"`
	Color        string             `json:"color"`
	VehicleType  helpers.EnumEntity `json:"vehicle_type"`
	OwnerId      int64              `json:"owner_id"`
}

func NewVehicle(licensePlate int64,
	brand, color string, vehicleType helpers.EnumEntity,
	ownerId int64) *Vehicle {
	return &Vehicle{
		LicensePlate: licensePlate,
		Brand:        brand,
		Color:        color,
		VehicleType:  vehicleType,
		OwnerId:      ownerId,
	}
}
