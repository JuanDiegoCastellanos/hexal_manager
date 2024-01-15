package entities

type Module struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Operations []Operation `json:"operations"`
}
type IModule interface{}

func NewModule(id int64, name string) *Module {
	return &Module{ID: id, Name: name}
}

func (m *Module) AddOperation(op Operation) []Operation {
	m.Operations = append(m.Operations, op)
	return m.Operations
}
