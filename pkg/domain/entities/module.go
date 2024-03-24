package entities

type Module struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Operations []int64 `json:"operations"`
}
