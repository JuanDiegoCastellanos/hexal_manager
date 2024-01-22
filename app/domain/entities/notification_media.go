package entities

type NotificationMedia struct {
	ID     int64
	Name   string
	Detail string
}

func NewNotificationMedia(id int64, name, detail string) *NotificationMedia {
	return &NotificationMedia{
		ID:     id,
		Name:   name,
		Detail: detail,
	}
}
