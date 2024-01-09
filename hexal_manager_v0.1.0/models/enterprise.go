package models

import "time"

type Enterprise struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Domain          string    `json:"domain"`
	SuscriptionCode string    `json:"suscription_code"`
	CreationDate    time.Time `json:"creation_date"`
}
