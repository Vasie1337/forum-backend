package services

import (
	"server/internal/models"
	"server/internal/repository"
)

type AdminService struct {
	UserRepo repository.UserRepository
	KeyRepo  repository.KeyRepository
}

func NewAdminService(userRepo repository.UserRepository, keyRepo repository.KeyRepository) *AdminService {
	return &AdminService{
		UserRepo: userRepo,
		KeyRepo:  keyRepo,
	}
}

func (s *AdminService) GetAllUsers() ([]*models.User, error) {
	return s.UserRepo.GetAll()
}

func (s *AdminService) GetUserByID(id int) (*models.User, error) {
	return s.UserRepo.GetByID(id)
}

func (s *AdminService) CreateUser(user *models.User) error {
	return s.UserRepo.Create(user)
}

func (s *AdminService) GetKeyByID(id int) (*models.Key, error) {
	return s.KeyRepo.GetByID(id)
}

func (s *AdminService) CreateKey(key *models.Key) error {
	return s.KeyRepo.Create(key)
}
