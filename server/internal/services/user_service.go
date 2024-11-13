package services

import (
	"server/internal/config"
	"server/internal/repository"
	"time"
)

type UserService struct {
	UserRepo repository.UserRepository
	KeyRepo  repository.KeyRepository
}

func NewUserService(userRepo repository.UserRepository, keyRepo repository.KeyRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
		KeyRepo:  keyRepo,
	}
}

func (s *UserService) RedeemKey(userID int, keyValue string) error {
	key, err := s.KeyRepo.GetByValue(keyValue)
	if err != nil {
		return err
	}

	if key.Redeemed {
		return config.ErrKeyAlreadyRedeemed
	}

	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return err
	}

	key.Redeemed = true
	key.RedeemedAt = time.Now()
	key.UserID = user.ID
	key.ValidUntil = time.Now().AddDate(0, 0, 7) // 7 days from now

	err = s.KeyRepo.Update(key)
	if err != nil {
		return err
	}

	return nil
}
