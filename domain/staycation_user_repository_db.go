package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"staycation/errs"
	"staycation/logger"
	"time"
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
	customerSql := "SELECT id, name, city, zipcode, dateOfBirth, status from StaycationUsers WHERE id = ?"

	var u StaycationUser

	err := d.client.Get(&u, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewInternalServerError("Unexpected database error")
		}
	}

	return &u, nil
}

func NewStaycationRepositoryDb() StaycationUserRepositoryDb {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, dbName)

	db, err := sqlx.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return StaycationUserRepositoryDb{client: db}
}
