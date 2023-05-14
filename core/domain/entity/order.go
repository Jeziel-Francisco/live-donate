package coredomainentity

import commonerror "github.com/jeziel-francisco/live-donate/common/error"

type OrderEntity struct {
	Amount     float64 `json:"amount"`
	Message    string  `json:"message"`
	ReceiverId string  `json:"receiver_id"`
}

func (o *OrderEntity) Validate() commonerror.ValidateError {
	if o.Amount <= 0 {
		return commonerror.NewValidateError("invalid amount", "order", 400)
	}

	if len(o.ReceiverId) <= 0 {
		return commonerror.NewValidateError("non-existent receiver", "order", 404)
	}

	return nil
}
