package coreusecaseorder

import (
	coreadapter "github.com/jeziel-francisco/live-donate/core/adapter"
	corecontrollerdto "github.com/jeziel-francisco/live-donate/core/controller/dto"
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	usecaseinterface "github.com/jeziel-francisco/live-donate/core/usecase/interface"
)

const (
	pubCreateOrderTopicName = "create_order_event"
)

func NewCreateOrderUseCase(orderRestRepository usecaseinterface.OrderRestRepository,
	orderDatabaseRepository usecaseinterface.OrderDatabaseRepository, messengerRepository usecaseinterface.MessengerRepository) *createOrder {
	return &createOrder{
		orderRestRepository:     orderRestRepository,
		orderDatabaseRepository: orderDatabaseRepository,
		messengerRepository:     messengerRepository,
	}
}

type createOrder struct {
	orderRestRepository     usecaseinterface.OrderRestRepository
	orderDatabaseRepository usecaseinterface.OrderDatabaseRepository
	messengerRepository     usecaseinterface.MessengerRepository
}

func (c *createOrder) Execute(entity coredomainentity.OrderEntity) (*corecontrollerdto.OutputCreateOrder, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	entityResult, err := c.orderRestRepository.Create(coreadapter.EntityOrderToRestRepositoryOrder(entity))
	if err != nil {
		return nil, err
	}
	entity.SetQrCode(entityResult.QrCode)
	entity.SetExternalID(entityResult.ExternalID)
	entity.SetTransactionID(entityResult.TransactionID)

	entityResult, err = c.orderDatabaseRepository.Save(coreadapter.EntityOrderToDatabaseRepositoryOrder(entity))
	if err != nil {
		return nil, err
	}
	entity.SetID(entityResult.ID)

	err = c.messengerRepository.SendTopic(entity, pubCreateOrderTopicName)
	if err != nil {
		return nil, err
	}

	return coreadapter.EntityOrderToControllerOrder(entity), nil
}
