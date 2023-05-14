package usecaseadapterrepository

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	repositorydto "github.com/jeziel-francisco/live-donate/repository/dto"
)

func EntityOrderToCreateOrderRestApiDto(entity coredomainentity.OrderEntity) repositorydto.CreateOrderRestApiDto {
	return repositorydto.CreateOrderRestApiDto{ReceiverId: entity.ReceiverId, TotalAmount: entity.Amount}
}
