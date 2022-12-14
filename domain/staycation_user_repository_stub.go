package domain

import "github.com/PatrickDei/log-lib/errs"

type StaycationUserRepositoryStub struct {
	staycationUsers []StaycationUser
}

func (s StaycationUserRepositoryStub) FindAll() ([]StaycationUser, *errs.AppError) {
	return s.staycationUsers, nil
}

func NewStaycationRepositoryStub() StaycationUserRepositoryStub {
	users := []StaycationUser{
		{
			Id:          "1",
			Name:        "John",
			City:        "Zagreb",
			Zipcode:     "10000",
			DateOfBirth: "2020-01-01",
			Status:      "1",
		},
		{
			Id:          "2",
			Name:        "Doe",
			City:        "Split",
			Zipcode:     "10002",
			DateOfBirth: "2020-01-01",
			Status:      "1",
		},
	}

	return StaycationUserRepositoryStub{staycationUsers: users}
}
