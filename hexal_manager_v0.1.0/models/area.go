package models

type Area struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}
