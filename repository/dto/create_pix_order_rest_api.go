package repositorydto

type CreateOrderRestApiDto struct {
	TotalAmount float64 `json:"total_amount"`
	ReceiverId  string  `json:"receiver_id"`
}
