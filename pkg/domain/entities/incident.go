package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Incident struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	State        helpers.EnumEntity `json:"state"`
	Description  string             `json:"description"`
	OwnerID      int64              `json:"owner_id"`
	Priority     helpers.EnumEntity `json:"priority"`
	TypeIncident helpers.EnumEntity `json:"type_incident"`
	CreatedAt    time.Time          `json:"created_at"`
}

func NewIncident(id int64, name, description string,
	ownerID int64, state, priority,
	typeService helpers.EnumEntity, createdAt time.Time) *Incident {
	return &Incident{
		ID:           id,
		Name:         name,
		State:        state,
		Description:  description,
		OwnerID:      ownerID,
		Priority:     priority,
		TypeIncident: typeService,
		CreatedAt:    createdAt,
	}
}
