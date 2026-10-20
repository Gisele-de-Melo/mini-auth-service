package service

import (
	"mini-auth-service/internal/model"
	"mini-auth-service/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetByEmail(email string) (*model.User, error) {
	return s.userRepo.FindByEmail(email)
}
