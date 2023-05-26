package testmockinfrarepositoryrest

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func NewOrderRestMock(successCreate coredomainentity.OrderEntity, errorCreate error) *OrderRestMock {
	return &OrderRestMock{succesCreate: successCreate, errorCreate: errorCreate}
}

type OrderRestMock struct {
	succesCreate coredomainentity.OrderEntity
	errorCreate  error
}

func (rest *OrderRestMock) Create(order infrarepositorydto.InputCreateOrderRestApiDto) (coredomainentity.OrderEntity, error) {
	return rest.succesCreate, rest.errorCreate
}
