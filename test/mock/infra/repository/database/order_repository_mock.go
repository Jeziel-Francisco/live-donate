package testmockinfrarepositorydatabase

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderDatabaseMock(saveError error) *OrderDatabaseMock {
	return &OrderDatabaseMock{saveError: saveError}
}

type OrderDatabaseMock struct {
	saveError error
}

func (rest *OrderDatabaseMock) Save(order infrarepositorydto.InputSaveOrderDatabaseDto) error {
	return rest.saveError
}
