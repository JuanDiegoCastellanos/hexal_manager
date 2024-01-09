package models

import "time"

type Project struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	State        bool      `json:"state"`
	LeaderName   string    `json:"leader_name"`
	AreaId       string    `json:"area_id"`
	EnterpriseId string    `json:"enterprise_id"`
}
