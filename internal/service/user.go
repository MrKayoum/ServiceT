package service

import (
	"ServiceT/internal/model"
	"ServiceT/internal/repository"
	"errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) RegisterUser(user *model.User) error {
	// Здесь можно добавить логику, например, хеширование пароля
	return s.userRepo.Create(user)
}

func (s *UserService) LoginUser(username, password string) (*model.User, error) {
	// Проверка пользователя с паролем
	user, err := s.userRepo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
