package entities

type NotificationMedia struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func NewNotificationMedia(id int64, name, detail string) *NotificationMedia {
	return &NotificationMedia{
		ID:     id,
		Name:   name,
		Detail: detail,
	}
}
