package coreusecaseorder

import (
	"testing"

	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	coreusecaseorder "github.com/jeziel-francisco/live-donate/core/usecase/order"
	testmockinfrarepositorydatabase "github.com/jeziel-francisco/live-donate/test/mock/infra/repository/database"
	testmockinfrarepositoryrest "github.com/jeziel-francisco/live-donate/test/mock/infra/repository/rest"
)

func TestExecuteInvalidOrder(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 0, ReceiverId: "1"}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock()
	orderDatabaseMock := testmockinfrarepositorydatabase.NewOrderDatabaseMock()
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, orderDatabaseMock)
	_, err := createOrderUseCase.Execute(order)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecute(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100, ReceiverId: "1"}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock()
	orderDatabaseMock := testmockinfrarepositorydatabase.NewOrderDatabaseMock()
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, orderDatabaseMock)
	output, err := createOrderUseCase.Execute(order)
	if err != nil {
		t.Errorf("unexpected error")
	}

	if len(output.QrCode) == 0 {
		t.Errorf("expected qrcode data")
	}
}
