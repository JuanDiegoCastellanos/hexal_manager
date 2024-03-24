package entities

type ParkingLot struct {
	ID                   int64 `json:"id"`
	ResidentialComplexID int64 `json:"residential_complex_id"`
	LargeSpaceCount      int64 `json:"large_space_count"`
	MediumSpaceCount     int64 `json:"medium_space_count"`
	SmallSpaceCount      int64 `json:"small_space_count"`
}
