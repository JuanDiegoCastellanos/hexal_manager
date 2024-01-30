package entities

import "time"

type RealUser struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastname"`
	Document    string    `json:"document"`
	PhoneNumber string    `json:"phone_number"`
	Occupation  string    `json:"occupation"`
	PetsCount   int       `json:"pets_count"`
	ApartmentId int64     `json:"apartment_id"`
	VehicleId   int64     `json:"vehicle_id"`
	LeaseId     int64     `json:"lease_id"`
	UserId      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewRealUser(id int64, name, lastName, document, phoneNumber,
	occupation string, petsCount int, apartmentId, vehicleId, leaseId, userId int64, createdAt time.Time) *RealUser {
	return &RealUser{
		ID:          id,
		Name:        name,
		LastName:    lastName,
		Document:    document,
		PhoneNumber: phoneNumber,
		Occupation:  occupation,
		PetsCount:   petsCount,
		ApartmentId: apartmentId,
		VehicleId:   vehicleId,
		LeaseId:     leaseId,
		UserId:      userId,
		CreatedAt:   createdAt,
	}
}
