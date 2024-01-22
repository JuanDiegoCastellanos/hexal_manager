package entities

type PaymentMethod struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

func NewPaymentMethod(id int64, name string, provider string) *PaymentMethod {
	return &PaymentMethod{
		ID:       id,
		Name:     name,
		Provider: provider,
	}
}
