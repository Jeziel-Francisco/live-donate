package coreadapter

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func EntityOrderToRestRepositoryOrder(entity coredomainentity.OrderEntity) infrarepositorydto.InputCreateOrderRestApiDto {
	return infrarepositorydto.InputCreateOrderRestApiDto{ReceiverID: entity.ReceiverID, TotalAmount: entity.Amount}
}

func EntityOrderToDatabaseRepositoryOrder(entity coredomainentity.OrderEntity) infrarepositorydto.InputSaveOrderDatabaseDto {
	return infrarepositorydto.InputSaveOrderDatabaseDto{Amount: entity.Amount, Message: entity.Message, ReceiverID: entity.ReceiverID, QrCode: entity.QrCode}
}
