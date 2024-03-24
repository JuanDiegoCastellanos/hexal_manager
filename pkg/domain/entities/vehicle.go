package entities

import "hexal_manager/util/helpers"

type Vehicle struct {
	LicensePlate int64              `json:"license_plate"`
	Brand        string             `json:"brand"`
	Color        string             `json:"color"`
	VehicleType  helpers.EnumEntity `json:"vehicle_type"`
	OwnerId      int64              `json:"owner_id"`
}
