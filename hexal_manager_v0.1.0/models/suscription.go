package models

import "time"

type Suscription struct {
	SuscriptionCode   string    `json:"suscription_code"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	IsValid           bool      `json:"is_valid"`
	SuscriptionTypeId string    `json:"suscription_type_id"`
}
