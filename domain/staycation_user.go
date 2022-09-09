package domain

import (
	"github.com/PatrickDei/log-lib/errs"
	"staycation/dto"
)

type StaycationUser struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"dateOfBirth"`
	Status      string
}

func (u StaycationUser) statusAsText() string {
	statusAsText := "Active"

	if u.Status == "0" {
		statusAsText = "Inactive"
	}

	return statusAsText
}

func (u StaycationUser) ToDto() dto.StaycationUserResponse {
	return dto.StaycationUserResponse{
		Id:          u.Id,
		Name:        u.Name,
		City:        u.City,
		Zipcode:     u.Zipcode,
		DateOfBirth: u.DateOfBirth,
		Status:      u.statusAsText(),
	}
}

type StaycationUserRepository interface {
	FindAll() ([]StaycationUser, *errs.AppError)
	FindById(string) (*StaycationUser, *errs.AppError)
}
