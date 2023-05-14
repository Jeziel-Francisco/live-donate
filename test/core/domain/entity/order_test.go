package testcoredomainentity

import (
	"testing"

	coredomainentity "github.com/jeziel-francisco/live-donate/core/domain/entity"
)

func TestValidateInvalidAmount(t *testing.T) {
	order := coredomainentity.OrderEntity{}

	err := order.Validate()

	if err == nil {
		t.Errorf("Expected err")
	}

	if err.Error() != "invalid amount" {
		t.Errorf("invalid message")
	}

	if err.Entity() != "order" {
		t.Errorf("invalid entity")
	}

	if err.Code() != 400 {
		t.Errorf("invalid code")
	}
}

func TestValidateInvalidReceiverId(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100}

	err := order.Validate()

	if err == nil {
		t.Errorf("Expected err")
	}

	if err.Error() != "non-existent receiver" {
		t.Errorf("invalid message")
	}

	if err.Entity() != "order" {
		t.Errorf("invalid entity")
	}

	if err.Code() != 404 {
		t.Errorf("invalid code")
	}
}

func TestValidate(t *testing.T) {
	order := coredomainentity.OrderEntity{Amount: 100, ReceiverId: "100", Message: "test"}

	err := order.Validate()

	if err != nil {
		t.Errorf("Unexpected err")
	}
}
