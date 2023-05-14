package coreusecaseorder

import (
	usecaseinterface "github.com/jeziel-francisco/live-donate/core/usecase/interface"
)

func NewCreateOrderUseCase(orderRepository usecaseinterface.OrderRepository) *createOrder {
	return &createOrder{
		orderRepository: orderRepository,
	}
}
