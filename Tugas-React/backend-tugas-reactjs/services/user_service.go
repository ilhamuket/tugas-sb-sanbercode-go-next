package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type UserService interface {
	CreateUser(input models.CreateUserInput) (models.User, error)
	UpdateUser(id uint, input models.UpdateUserInput) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input models.CreateUserInput) (models.User, error) {
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	newUser, err := s.userRepository.CreateUser(user)
	return newUser, err
}

func (s *userService) UpdateUser(id uint, input models.UpdateUserInput) (models.User, error) {
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return user, err
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		user.Password = input.Password
	}

	updatedUser, err := s.userRepository.UpdateUser(user)
	return updatedUser, err
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepository.FindUserByID(id)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return err
	}

	return s.userRepository.DeleteUser(user)
}
