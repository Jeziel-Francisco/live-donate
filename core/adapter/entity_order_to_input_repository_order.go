package coreadapter

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func EntityOrderToRestRepositoryOrder(entity coredomainentity.OrderEntity) infrarepositorydto.InputCreateOrderRestApiDto {
	return infrarepositorydto.InputCreateOrderRestApiDto{ReceiverId: entity.ReceiverId, TotalAmount: entity.Amount}
}

func EntityOrderToDatabaseRepositoryOrder(entity coredomainentity.OrderEntity) infrarepositorydto.InputSaveOrderDatabaseDto {
	return infrarepositorydto.InputSaveOrderDatabaseDto{Amount: entity.Amount, Message: entity.Message, ReceiverId: entity.ReceiverId, QrCode: entity.QrCode}
}
