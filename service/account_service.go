package service

import (
	"github.com/PatrickDei/log-lib/errs"
	"staycation/domain"
	"staycation/dto"
	"time"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (as DefaultAccountService) NewAccount(req dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		UserId:      req.UserId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	account, err := as.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := account.ToResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
