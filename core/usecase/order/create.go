package coreusecaseorder

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	usecaseadapterrepository "github.com/jeziel-francisco/live-donate/core/usecase/adapter/repository"
	usecaseinterface "github.com/jeziel-francisco/live-donate/core/usecase/interface"
)

type createOrder struct {
	orderRepository usecaseinterface.OrderRepository
}

func (c *createOrder) Execute(entity coredomainentity.OrderEntity) error {
	err := entity.Validate()
	if err != nil {
		return err
	}

	return c.orderRepository.Create(usecaseadapterrepository.EntityOrderToCreateOrderRestApiDto(entity))
}
