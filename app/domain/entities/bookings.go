package entities

import (
	"hexal_manager/util/helpers"
	"time"
)

type Booking struct {
	ID                 int64              `json:"id"`
	SocialSpaceID      int64              `json:"social_space_id"`
	StartDate          time.Time          `json:"start_date"`
	EndDate            time.Time          `json:"end_date"`
	Status             helpers.EnumEntity `json:"status"`
	NumberOfPeople     int64              `json:"number_of_people"`
	SpecialRequest     string             `json:"special_request"`
	TotalCost          float64            `json:"total_cost"`
	PaymentStatus      string             `json:"payment_status"`
	PaymentMethodID    int64              `json:"payment_method_id"`
	CancellationReason string             `json:"cancellation_reason"`
	UserID             int64              `json:"user_id"`
	CreatedAt          time.Time          `json:"created_at"`
}

func NewBooking(id, socialSpace int64,
	startDate, endDate time.Time, status helpers.EnumEntity, numberOfPeople int64,
	specialRequest string, totalCost float64, paymentStatus string, paymentMethodID int64,
	cancellationReason string, userID int64, createdAt time.Time) *Booking {

	return &Booking{
		ID:                 id,
		SocialSpaceID:      socialSpace,
		StartDate:          startDate,
		EndDate:            endDate,
		Status:             status,
		NumberOfPeople:     numberOfPeople,
		SpecialRequest:     specialRequest,
		TotalCost:          totalCost,
		PaymentStatus:      paymentStatus,
		PaymentMethodID:    paymentMethodID,
		CancellationReason: cancellationReason,
		UserID:             userID,
		CreatedAt:          createdAt,
	}
}
