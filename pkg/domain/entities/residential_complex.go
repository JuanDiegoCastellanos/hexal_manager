package entities

type ResidentialComplex struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	ApartmentsCount int64  `json:"apartments_count"`
}
