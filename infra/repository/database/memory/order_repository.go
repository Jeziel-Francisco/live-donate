package infrarepositorymemory

import (
	coreadapter "github.com/jeziel-francisco/live-donate/core/adapter"
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func NewOrderDatabaseMemory() *OrderDatabaseMemory {
	return &OrderDatabaseMemory{}
}

type OrderDatabaseMemory struct {
}

func (rest *OrderDatabaseMemory) Save(order infrarepositorydto.InputSaveOrderDatabaseDto) (coredomainentity.OrderEntity, error) {
	return coreadapter.DatabaseRepositorySaveOrderToEntityOrder(infrarepositorydto.OutputSaveOrderDatabaseDto{}), nil
}
