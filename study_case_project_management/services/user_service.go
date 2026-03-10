package services

import (
	"errors"

	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/repositories"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
} // hanya bisa diakses di package ini aja

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService)Register(user *models.User) error{
	// cek email yg terdaftar apakah sudah dipake atau blm
	// hashing password
	// set role
	// simpan user

	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()
	return s.repo.Create(user)
}