package dto

import (
	"net/http"
	"testing"
)

func TestShouldReturnErrorOnWrongAccountType(t *testing.T) {
	// Arrange
	request := AccountRequest{AccountType: "Invalid transaction type", Amount: 55555}
	// Act
	appError := request.Validate()
	// Assert
	if appError.Message != "Account type should be checking or savings" {
		t.Error("Invalid message on wrong account type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code returned")
	}
}

func TestShouldReturnErrorOnWrongAmount(t *testing.T) {
	// Arrange
	request := AccountRequest{AccountType: "savings", Amount: 2}
	// Act
	appError := request.Validate()
	// Assert
	if appError.Message != "To open a new account you need to deposit at least 5000" {
		t.Error("Invalid message on wrong account amount")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code returned")
	}
}
