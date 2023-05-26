package usecaseinterface

import (
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"
)

type OrderRestRepository interface {
	Create(order infrarepositorydto.InputCreateOrderRestApiDto) (coredomainentity.OrderEntity, error)
}
