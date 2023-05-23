package testmockinfrarepositorydatabase

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderDatabaseMock() *OrderDatabaseMock {
	return &OrderDatabaseMock{}
}

type OrderDatabaseMock struct {
}

func (rest *OrderDatabaseMock) Save(order infrarepositorydto.InputSaveOrderDatabaseDto) error {
	return nil
}
