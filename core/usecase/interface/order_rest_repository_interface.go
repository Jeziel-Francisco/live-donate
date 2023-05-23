package usecaseinterface

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

type OrderRestRepository interface {
	Create(order infrarepositorydto.InputCreateOrderRestApiDto) (string, error)
}
