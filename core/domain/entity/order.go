package coredomainentity

import commonerror "github.com/jeziel-francisco/live-donate/common/error"

type OrderEntity struct {
	Amount     float64 `json:"amount"`
	Message    string  `json:"message"`
	ReceiverId string  `json:"receiver_id"`
	QrCode     string  `json:"qr_code"`
}

func (o *OrderEntity) Validate() error {
	if o.Amount <= 0 {
		return commonerror.NewBusinessError("invalid amount", "order", 400)
	}

	if len(o.ReceiverId) <= 0 {
		return commonerror.NewBusinessError("non-existent receiver", "order", 404)
	}

	return nil
}

func (o *OrderEntity) SetQrCode(qrCode string) {
	o.QrCode = qrCode
}
