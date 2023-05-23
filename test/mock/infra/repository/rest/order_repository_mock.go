package testmockinfrarepositoryrest

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderRestMock() *OrderRestMock {
	return &OrderRestMock{}
}

type OrderRestMock struct {
}

func (rest *OrderRestMock) Create(order infrarepositorydto.InputCreateOrderRestApiDto) (string, error) {
	return "", nil
}
