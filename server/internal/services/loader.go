package services

import (
	"server/internal/models"
	"server/internal/repository"
)

type LoaderService struct {
	KeyRepo repository.KeyRepository
}

func NewLoaderService(keyRepo repository.KeyRepository) *LoaderService {
	return &LoaderService{
		KeyRepo: keyRepo,
	}
}

func (s *LoaderService) GetCheatData(userkey *models.Key) (*models.Cheat, error) {
	//cheat, err := s.KeyRepo.GetCheatData(userkey)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
