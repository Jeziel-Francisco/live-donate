package coreusecaseorder

import (
	"testing"

	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	coreusecaseorder "github.com/jeziel-francisco/live-donate/core/usecase/order"
	repositoryrest "github.com/jeziel-francisco/live-donate/repository/rest"
)

func TestExecuteInvalidOrder(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 0, ReceiverId: "1"}
	orderRestApi := repositoryrest.NewOrderRestApi()
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestApi)
	err := createOrderUseCase.Execute(order)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecute(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100, ReceiverId: "1"}
	orderRestApi := repositoryrest.NewOrderRestApi()
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestApi)
	err := createOrderUseCase.Execute(order)
	if err != nil {
		t.Errorf("unexpected error")
	}
}
