package dto

import (
	"github.com/PatrickDei/log-lib/errs"
	"strings"
)

type AccountRequest struct {
	UserId      string
	AccountType string
	Amount      float64
}

func (ar AccountRequest) Validate() *errs.AppError {
	if ar.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}
	if strings.ToLower(ar.AccountType) != "savings" && strings.ToLower(ar.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or savings")
	}
	return nil
}
