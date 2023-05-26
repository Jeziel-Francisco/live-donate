package coredomainentity

import commonerror "github.com/jeziel-francisco/live-donate/common/error"

type OrderEntity struct {
	Amount        float64
	Message       string
	ReceiverID    string
	QrCode        string
	ID            string
	ExternalID    string
	TransactionID string
}

func (o *OrderEntity) Validate() error {
	if o.Amount <= 0 {
		return commonerror.NewBusinessError("invalid amount", "order", 400)
	}

	if len(o.ReceiverID) <= 0 {
		return commonerror.NewBusinessError("non-existent receiver", "order", 404)
	}

	return nil
}

func (o *OrderEntity) SetID(ID string) {
	o.ID = ID
}

func (o *OrderEntity) SetQrCode(qrCode string) {
	o.QrCode = qrCode
}

func (o *OrderEntity) SetExternalID(externalID string) {
	o.ExternalID = externalID
}

func (o *OrderEntity) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}
