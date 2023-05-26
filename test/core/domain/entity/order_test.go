package testcoredomainentity

import (
	"testing"

	commonerror "github.com/jeziel-francisco/live-donate/common/error"
	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
)

func TestValidateInvalidAmount(t *testing.T) {
	order := coredomainentity.OrderEntity{}

	err := order.Validate()

	if err == nil {
		t.Errorf("Expected err")
	}

	businessError := err.(commonerror.BusinessError)

	if businessError.Error() != "invalid amount" {
		t.Errorf("invalid message")
	}

	if businessError.Entity() != "order" {
		t.Errorf("invalid entity")
	}

	if businessError.Code() != 400 {
		t.Errorf("invalid code")
	}
}

func TestValidateInvalidReceiverID(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100}

	err := order.Validate()

	if err == nil {
		t.Errorf("Expected err")
	}

	businessError := err.(commonerror.BusinessError)

	if businessError.Error() != "non-existent receiver" {
		t.Errorf("invalid message")
	}

	if businessError.Entity() != "order" {
		t.Errorf("invalid entity")
	}

	if businessError.Code() != 404 {
		t.Errorf("invalid code")
	}
}

func TestValidate(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100, ReceiverID: "100", Message: "test"}

	err := order.Validate()

	if err != nil {
		t.Errorf("Unexpected err")
	}
}
