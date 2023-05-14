package testcommonerror

import (
	"testing"

	commonerror "github.com/jeziel-francisco/live-donate/common/error"
)

func TestNewValidateError(t *testing.T) {
	messageError := "test"
	entityError := "any"
	var codeError int64 = 400
	err := commonerror.NewValidateError(messageError, entityError, codeError)

	if err == nil {
		t.Errorf("Expected err")
	}

	if err.Error() != messageError {
		t.Errorf("Invalid message error")
	}

	if err.Entity() != entityError {
		t.Errorf("Invalid entity error")
	}

	if err.Code() != codeError {
		t.Errorf("Invalid code error")
	}
}
