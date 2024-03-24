package entities

type Rol struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Operations []int64 `json:"operations"`
}
