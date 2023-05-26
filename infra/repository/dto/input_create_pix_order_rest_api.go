package infrarepositorydto

type InputCreateOrderRestApiDto struct {
	TotalAmount float64 `json:"total_amount"`
	ReceiverID  string  `json:"receiver_id"`
}
