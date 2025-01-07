package services

import (
	"errors"
	"time"

	"github.com/pvfarooq/go-gin-crud/helpers"
	"github.com/pvfarooq/go-gin-crud/models"
	"github.com/pvfarooq/go-gin-crud/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	existingUser, err := s.userRepo.FetchUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email address already taken")
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.CreateUser(user)
}

func (s *UserService) FetchUsers() (*[]models.User, error) {
	return s.userRepo.FetchUsers()
}

func (s *UserService) FetchUser(id string) (*models.User, error) {
	return s.userRepo.FetchUser(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	user.UpdatedAt = time.Now()
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.userRepo.DeleteUser(user)
}
