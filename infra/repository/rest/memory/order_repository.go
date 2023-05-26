package infrarepositoryrestmemory

import (
	coreadapter "github.com/jeziel-francisco/live-donate/core/adapter"
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func NewOrderRestMemory() *OrderRestMemory {
	return &OrderRestMemory{}
}

type OrderRestMemory struct {
}

func (rest *OrderRestMemory) Create(order infrarepositorydto.InputCreateOrderRestApiDto) (coredomainentity.OrderEntity, error) {
	return coreadapter.RestRepositoryCreateOrderToEntityOrder(infrarepositorydto.OutputCreateOrderRestApiDto{}), nil
}
