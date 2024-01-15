package entities

type Rol struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Operations []Operation `json:"operations"`
}

func NewRol(id int64, name string, operations []Operation) *Rol {
	return &Rol{
		ID:         id,
		Name:       name,
		Operations: operations,
	}
}

func (r *Rol) AddOperation(op Operation) *[]Operation {
	r.Operations = append(r.Operations, op)
	return &r.Operations
}
