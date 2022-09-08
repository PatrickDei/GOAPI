package service

import (
	"staycation/domain"
	"staycation/errs"
)

type StaycationUserService interface {
	GetAllStaycationUsers() ([]domain.StaycationUser, error)
	GetStaycationUserById(string) (*domain.StaycationUser, *errs.AppError)
}

type DefaultStaycationUserService struct {
	repo domain.StaycationUserRepository
}

func (s DefaultStaycationUserService) GetAllStaycationUsers() ([]domain.StaycationUser, error) {
	return s.repo.FindAll()
}

func (s DefaultStaycationUserService) GetStaycationUserById(id string) (*domain.StaycationUser, *errs.AppError) {
	return s.repo.FindById(id)
}

func NewStaycationUserService(repository domain.StaycationUserRepository) DefaultStaycationUserService {
	return DefaultStaycationUserService{repo: repository}
}
