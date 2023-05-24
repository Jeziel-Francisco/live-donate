package testmockinfrarepositoryrest

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

func NewOrderRestMock(successCreate string, errorCreate error) *OrderRestMock {
	return &OrderRestMock{succesCreate: successCreate, errorCreate: errorCreate}
}

type OrderRestMock struct {
	succesCreate string
	errorCreate  error
}

func (rest *OrderRestMock) Create(order infrarepositorydto.InputCreateOrderRestApiDto) (string, error) {
	return rest.succesCreate, rest.errorCreate
}
