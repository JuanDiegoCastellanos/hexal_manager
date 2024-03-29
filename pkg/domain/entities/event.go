package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Event struct {
	ID                  int64              `json:"id"`
	Name                string             `json:"name"`
	Priority            helpers.EnumEntity `json:"priority"`
	TypeEvent           helpers.EnumEntity `json:"type_event"`
	Description         string             `json:"description"`
	StartDate           time.Time          `json:"start_date"`
	CreatedAt           time.Time          `json:"created_at"`
	SocialSpacesID      int64              `json:"social_spaces_id"`
	OwnerID             int64              `json:"owner_id"`
	NotificationMediaID int64              `json:"notification_media_id"`
}

func NewEvent(id int64, name string, priority, typeEvent helpers.EnumEntity,
	description string, startDate, createdAt time.Time, socialSpacesID, ownerID, notificationMediaID int64) *Event {

	return &Event{
		ID:                  id,
		Name:                name,
		Priority:            priority,
		TypeEvent:           typeEvent,
		Description:         description,
		StartDate:           startDate,
		CreatedAt:           createdAt,
		SocialSpacesID:      socialSpacesID,
		OwnerID:             ownerID,
		NotificationMediaID: notificationMediaID}
}
