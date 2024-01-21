package entities

import "time"

type Lease struct {
	ID              int64     `json:"id"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	TotalCost       float64   `json:"total_cost"`
	PaymentMethodID int64     `json:"payment_method_id"`
}

func NewLease(id int64, startDate time.Time, endDate time.Time, totalCost float64, paymentMethod int64) *Lease {
	return &Lease{
		ID:              id,
		StartDate:       startDate,
		EndDate:         endDate,
		TotalCost:       totalCost,
		PaymentMethodID: paymentMethod,
	}
}
