package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Incident struct {
	ID          int64
	Name        string
	State       helpers.EnumEntity
	Description string
	OwnerID     int64
	Priority    helpers.EnumEntity
	Type        helpers.EnumEntity
	CreatedAt   time.Time
}

func NewIncident(id int64, name, description string,
	ownerID int64, state, priority,
	typeService helpers.EnumEntity, createdAt time.Time) *Incident {
	return &Incident{
		ID:          id,
		Name:        name,
		State:       state,
		Description: description,
		OwnerID:     ownerID,
		Priority:    priority,
		Type:        typeService,
		CreatedAt:   createdAt,
	}
}
