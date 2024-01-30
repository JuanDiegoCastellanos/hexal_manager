package entities

type Operation struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Roles []Rol  `json:"roles"`
}

func NewOperation(id int64, name string, roles []Rol) *Operation {
	return &Operation{ID: id, Name: name, Roles: roles}
}
func (o *Operation) AddRol(rol Rol) *[]Rol {
	o.Roles = append(o.Roles, rol)
	return &o.Roles
}
