package domain

import (
	"staycation/dto"
	"staycation/errs"
)

type Account struct {
	AccountId   string
	UserId      string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToResponseDto() dto.AccountResponse {
	return dto.AccountResponse{Id: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
