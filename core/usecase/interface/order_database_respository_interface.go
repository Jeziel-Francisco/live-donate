package usecaseinterface

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

type OrderDatabaseRepository interface {
	Save(order infrarepositorydto.InputSaveOrderDatabaseDto) (coredomainentity.OrderEntity, error)
}
