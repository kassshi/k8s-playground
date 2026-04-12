package service

import "github.com/kassshi/golang-practice/backend/internal/repository"

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}
