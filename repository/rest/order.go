package repositoryrest

import repositorydto "github.com/jeziel-francisco/live-donate/repository/dto"

type OrderRestApi struct {
}

func (rest *OrderRestApi) Create(order repositorydto.CreateOrderRestApiDto) error {
	return nil
}
