package services

import (
	"errors"
	"time"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(input models.CreateUserInput) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(id uint, input models.UpdateUserInput) (models.User, error)
	DeleteUser(id uint) error
	Login(input models.LoginInput) (models.User, string, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) CreateUser(input models.CreateUserInput) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	return s.repository.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *userService) UpdateUser(id uint, input models.UpdateUserInput) (models.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, err
		}
		user.Password = string(hashedPassword)
	}

	return s.repository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteUser(user)
}

func (s *userService) Login(input models.LoginInput) (models.User, string, error) {
	user, err := s.repository.GetUserByEmail(input.Email)
	if err != nil {
		return user, "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, "", errors.New("incorrect password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return user, "", err
	}

	return user, token, nil
}

func generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
