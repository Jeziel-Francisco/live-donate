package coreusecaseorder

import (
	"errors"
	"testing"

	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
	coreusecaseorder "github.com/jeziel-francisco/live-donate/core/usecase/order"
	testmockinfrarepositorydatabase "github.com/jeziel-francisco/live-donate/test/mock/infra/repository/database"
	testmockinfrarepositorypubsub "github.com/jeziel-francisco/live-donate/test/mock/infra/repository/pubsub"
	testmockinfrarepositoryrest "github.com/jeziel-francisco/live-donate/test/mock/infra/repository/rest"
)

func TestExecuteInvalidOrder(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 0, ReceiverID: "1"}
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(nil, nil, nil)
	_, err := createOrderUseCase.Execute(order)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteCreateOrderError(t *testing.T) {
	inputOrderUseCase := coredomainentity.OrderEntity{Amount: 0, ReceiverID: "1"}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock(coredomainentity.OrderEntity{}, errors.New("errror"))
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, nil, nil)
	_, err := createOrderUseCase.Execute(inputOrderUseCase)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteSaveOrderError(t *testing.T) {
	inputOrderUseCase := coredomainentity.OrderEntity{Amount: 100, ReceiverID: "1"}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock(coredomainentity.OrderEntity{}, nil)
	orderDatabaseMock := testmockinfrarepositorydatabase.NewOrderDatabaseMock(coredomainentity.OrderEntity{}, errors.New("error"))
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, orderDatabaseMock, nil)
	_, err := createOrderUseCase.Execute(inputOrderUseCase)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecuteSendTopicError(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100, ReceiverID: "1"}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock(coredomainentity.OrderEntity{}, nil)
	orderDatabaseMock := testmockinfrarepositorydatabase.NewOrderDatabaseMock(coredomainentity.OrderEntity{}, nil)
	pubSubMock := testmockinfrarepositorypubsub.NewMessageMock(errors.New(""), nil)
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, orderDatabaseMock, pubSubMock)
	_, err := createOrderUseCase.Execute(order)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestExecute(t *testing.T) {
	inputOrderUseCase := coredomainentity.OrderEntity{Amount: 100, ReceiverID: "1"}
	outputOrderRest := coredomainentity.OrderEntity{
		Amount:        100,
		ReceiverID:    "1",
		QrCode:        "Teste",
		ExternalID:    "1",
		TransactionID: "1",
	}
	outputOrderDatabase := coredomainentity.OrderEntity{
		Amount:        100,
		ReceiverID:    "1",
		Message:       "Teste",
		QrCode:        "Teste",
		ID:            "1",
		ExternalID:    "1",
		TransactionID: "1",
	}
	orderRestMock := testmockinfrarepositoryrest.NewOrderRestMock(outputOrderRest, nil)
	orderDatabaseMock := testmockinfrarepositorydatabase.NewOrderDatabaseMock(outputOrderDatabase, nil)
	pubSubMock := testmockinfrarepositorypubsub.NewMessageMock(nil, nil)
	createOrderUseCase := coreusecaseorder.NewCreateOrderUseCase(orderRestMock, orderDatabaseMock, pubSubMock)
	output, err := createOrderUseCase.Execute(inputOrderUseCase)
	if err != nil {
		t.Errorf("unexpected error")
	}

	if len(output.QrCode) == 0 {
		t.Errorf("expected qrcode data")
	}
}
