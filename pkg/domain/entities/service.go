package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Service struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	TotalCost    float64            `json:"total_cost"`
	GeneratedAt  time.Time          `json:"generated_at"`
	ExpiringDate time.Time          `json:"expiring_date"`
	ServiceType  helpers.EnumEntity `json:"service_type"`
	OwnerID      int64              `json:"owner_id"`
	ApartmentID  int64              `json:"apartment_id"`
}
