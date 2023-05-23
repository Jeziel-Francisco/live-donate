package usecaseinterface

import infrarepositorydto "github.com/jeziel-francisco/live-donate/infra/repository/dto"

type OrderDatabaseRespository interface {
	Save(order infrarepositorydto.InputSaveOrderDatabaseDto) error
}
