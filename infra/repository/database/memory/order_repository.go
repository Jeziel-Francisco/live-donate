package infrarespositorymemory

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderDatabaseMemory() *OrderDatabaseMemory {
	return &OrderDatabaseMemory{}
}

type OrderDatabaseMemory struct {
}

func (rest *OrderDatabaseMemory) Save(order infrarepositorydto.InputSaveOrderDatabaseDto) error {
	return nil
}
