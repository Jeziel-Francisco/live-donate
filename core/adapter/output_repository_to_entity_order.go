package coreadapter

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

func RestRepositoryCreateOrderToEntityOrder(entity infrarepositorydto.OutputCreateOrderRestApiDto) coredomainentity.OrderEntity {
	return coredomainentity.OrderEntity{
		Amount:        entity.TotalAmount,
		ReceiverID:    entity.ReceiverID,
		QrCode:        entity.QrCode,
		ExternalID:    entity.ExternalID,
		TransactionID: entity.TransactionID,
	}
}

func DatabaseRepositorySaveOrderToEntityOrder(entity infrarepositorydto.OutputSaveOrderDatabaseDto) coredomainentity.OrderEntity {
	return coredomainentity.OrderEntity{
		Amount:        entity.Amount,
		Message:       entity.Message,
		ReceiverID:    entity.ReceiverID,
		QrCode:        entity.QrCode,
		ID:            entity.ID,
		ExternalID:    entity.ExternalID,
		TransactionID: entity.TransactionID,
	}
}
