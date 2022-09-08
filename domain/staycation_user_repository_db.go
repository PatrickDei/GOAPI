package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"staycation/errs"
	"staycation/logger"
)

type StaycationUserRepositoryDb struct {
	client *sqlx.DB
}

func (d StaycationUserRepositoryDb) FindAll() ([]StaycationUser, *errs.AppError) {
	findAllSql := "SELECT id, name, city, zipcode, dateOfBirth, status from StaycationUsers"

	users := make([]StaycationUser, 0)

	err := d.client.Select(&users, findAllSql)
	if err != nil {
		logger.Error("Unexpected database error")
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	return users, nil
}

func (d StaycationUserRepositoryDb) FindById(id string) (*StaycationUser, *errs.AppError) {
	userSql := "SELECT id, name, city, zipcode, dateOfBirth, status from StaycationUsers WHERE id = ?"

	var u StaycationUser

	err := d.client.Get(&u, userSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			logger.Error("Error while scanning user" + err.Error())
			return nil, errs.NewInternalServerError("Unexpected database error")
		}
	}

	return &u, nil
}

func NewStaycationRepositoryDb(dbClient *sqlx.DB) StaycationUserRepositoryDb {
	return StaycationUserRepositoryDb{client: dbClient}
}
