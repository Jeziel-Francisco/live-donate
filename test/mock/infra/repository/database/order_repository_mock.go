package testmockinfrarepositorydatabase

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func NewOrderDatabaseMock(saveSuccess coredomainentity.OrderEntity, saveError error) *OrderDatabaseMock {
	return &OrderDatabaseMock{saveSuccess: saveSuccess, saveError: saveError}
}

type OrderDatabaseMock struct {
	saveSuccess coredomainentity.OrderEntity
	saveError   error
}

func (rest *OrderDatabaseMock) Save(order infrarepositorydto.InputSaveOrderDatabaseDto) (coredomainentity.OrderEntity, error) {
	return rest.saveSuccess, rest.saveError
}
