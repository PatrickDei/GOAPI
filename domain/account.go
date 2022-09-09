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

//go:generate mockgen -destination=../mocks/domain/mock_account_repository.go -package=domain staycation/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
