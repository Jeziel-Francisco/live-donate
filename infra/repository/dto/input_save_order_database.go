package infrarepositorydto

type InputSaveOrderDatabaseDto struct {
	ID         string  `json:"id"`
	Amount     float64 `json:"amount"`
	Message    string  `json:"message"`
	ReceiverID string  `json:"receiver_id"`
	QrCode     string  `json:"qr_code"`
}
