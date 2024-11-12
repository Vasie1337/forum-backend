package services

import (
	"server/internal/models"
	"server/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AdminRepo repository.AdminRepository
	UserRepo  repository.UserRepository
}

func (s *AuthService) AdminLogin(username, password string) (*models.Admin, error) {
	admin, err := s.AdminRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *AuthService) UserLogin(username, password string) (*models.User, error) {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
