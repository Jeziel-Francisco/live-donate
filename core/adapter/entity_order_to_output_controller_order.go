package coreadapter

import (
	corecontrollerdto "github.com/jeziel-francisco/live-donate/core/controller/dto"
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
)

func EntityOrderToControllerOrder(entity coredomainentity.OrderEntity) *corecontrollerdto.OutputCreateOrder {
	return &corecontrollerdto.OutputCreateOrder{}
}
