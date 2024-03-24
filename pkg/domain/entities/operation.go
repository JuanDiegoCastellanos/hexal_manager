package entities

type Operation struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Roles []int64 `json:"roles"`
}
