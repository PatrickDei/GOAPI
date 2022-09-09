package service

import (
	"github.com/PatrickDei/log-lib/errs"
	"github.com/golang/mock/gomock"
	realdomain "staycation/domain"
	"staycation/dto"
	"staycation/mocks/domain"
	"testing"
	"time"
)

func TestShouldReturnErrorWhenRequestIsNotValid(t *testing.T) {
	// Arrange
	request := dto.AccountRequest{
		UserId:      "1",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	// Act
	_, appError := service.NewAccount(request)

	// Assert
	if appError == nil {
		t.Error("Didn't throw error on invalid input")
	}
}

var mockRepository *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepository = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepository)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestShouldReturnErrorWhenAccountCannotBeCreated(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.AccountRequest{
		UserId:      "1",
		AccountType: "savings",
		Amount:      60000,
	}

	a := realdomain.Account{
		AccountId:   "",
		UserId:      req.UserId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepository.EXPECT().Save(a).Return(nil, errs.NewInternalServerError("Database error"))

	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("Test fail while validating")
	}
}

func TestReturnAccountWhenSaveIsSuccessful(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.AccountRequest{
		UserId:      "1",
		AccountType: "savings",
		Amount:      60000,
	}

	a := realdomain.Account{
		AccountId:   "",
		UserId:      req.UserId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	accountWithId := a
	accountWithId.AccountId = "1"

	mockRepository.EXPECT().Save(a).Return(&accountWithId, nil)

	// Act
	createdAccount, appError := service.NewAccount(req)

	// Assert
	if appError != nil {
		t.Error("Failed creating new account")
	}
	if createdAccount.Id != accountWithId.AccountId {
		t.Error("The saved Id doesn't match")
	}
}
