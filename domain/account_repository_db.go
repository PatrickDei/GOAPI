package domain

import (
	"github.com/jmoiron/sqlx"
	"staycation/errs"
	"staycation/logger"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (ar AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO Accounts (userId, openingDate, accountType, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := ar.client.Exec(sqlInsert, a.UserId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating your account: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
