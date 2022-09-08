package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"staycation/errs"
	"time"
)

type StaycationUserRepositoryDb struct {
	client *sql.DB
}

func (d StaycationUserRepositoryDb) FindAll() ([]StaycationUser, error) {
	findAllSql := "SELECT id, name, city, zipcode, dateOfBirth, status from StaycationUsers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error  fetchig data from db")
		return nil, err
	}

	users := make([]StaycationUser, 0)
	for rows.Next() {
		var u StaycationUser
		err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)

		if err != nil {
			log.Println("Error scanning users")
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (d StaycationUserRepositoryDb) FindById(id string) (*StaycationUser, *errs.AppError) {
	customerSql := "SELECT id, name, city, zipcode, dateOfBirth, status from StaycationUsers WHERE id = ?"

	row := d.client.QueryRow(customerSql, id)
	var u StaycationUser
	err := row.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer" + err.Error())
			return nil, errs.NewInternalServerError("Unexpected database error")
		}
	}

	return &u, nil
}

func NewStaycationRepositoryDb() StaycationUserRepositoryDb {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/Staycation")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return StaycationUserRepositoryDb{client: db}
}
