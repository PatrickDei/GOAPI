package service

import (
	"staycation/domain"
	"staycation/dto"
	"staycation/errs"
)

//go:generate mockgen -destination=../mocks/service/mock_staycation_user_service.go -package=service staycation/service StaycationUserService
type StaycationUserService interface {
	GetAllStaycationUsers() ([]dto.StaycationUserResponse, *errs.AppError)
	GetStaycationUserById(string) (*dto.StaycationUserResponse, *errs.AppError)
}

type DefaultStaycationUserService struct {
	repo domain.StaycationUserRepository
}

func (s DefaultStaycationUserService) GetAllStaycationUsers() ([]dto.StaycationUserResponse, *errs.AppError) {
	users, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.StaycationUserResponse, 0)
	for _, user := range users {
		response = append(response, user.ToDto())
	}

	return response, nil
}

func (s DefaultStaycationUserService) GetStaycationUserById(id string) (*dto.StaycationUserResponse, *errs.AppError) {
	u, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := u.ToDto()

	return &response, nil
}

func NewStaycationUserService(repository domain.StaycationUserRepository) DefaultStaycationUserService {
	return DefaultStaycationUserService{repo: repository}
}
