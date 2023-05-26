package infrarepositorydto

type OutputCreateOrderRestApiDto struct {
	TotalAmount   float64 `json:"total_amount"`
	ReceiverID    string  `json:"receiver_id"`
	QrCode        string  `json:"qr_code"`
	ExternalID    string  `json:"external_id"`
	TransactionID string  `json:"transaction_id"`
}
