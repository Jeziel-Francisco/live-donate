package infrarepositoryrestmemory

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderRestMemory() *OrderRestMemory {
	return &OrderRestMemory{}
}

type OrderRestMemory struct {
}

func (rest *OrderRestMemory) Create(order infrarepositorydto.InputCreateOrderRestApiDto) (string, error) {
	return "", nil
}
