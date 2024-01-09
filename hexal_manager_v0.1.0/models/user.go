package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Phone        string    `json:"phone"`
	Document     string    `json:"document"`
	DocumentType string    `json:"document_type"`
	BirthDate    time.Time `json:"birth_date"`
	Charge       string    `json:"charge"`
	EnterpriseId string    `json:"enterprise_id"`
	RolId        string    `json:"rol_id"`
}
