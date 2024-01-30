package entities

type Receipt struct {
	ID              int64   `json:"id"`
	OwnerID         int64   `json:"owner_id"`
	ServiceID       int64   `json:"service_id"`
	NotificationID  int64   `json:"notification_id"`
	PaymentMethodID int64   `json:"payment_method_id"`
	TotalCost       float64 `json:"total_cost"`
}

func NewReceipt(id, ownerID, serviceID, notificationID, paymentMethodID int64, totalCost float64) *Receipt {
	return &Receipt{
		ID:              id,
		OwnerID:         ownerID,
		ServiceID:       serviceID,
		NotificationID:  notificationID,
		PaymentMethodID: paymentMethodID,
		TotalCost:       totalCost,
	}
}
