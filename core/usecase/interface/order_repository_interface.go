package usecaseinterface

import repositorydto "github.com/jeziel-francisco/live-donate/repository/dto"

type OrderRepository interface {
	Create(order repositorydto.CreateOrderRestApiDto) error
}
