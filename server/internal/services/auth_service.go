package services

import (
	"errors"
	"server/internal/config"
	"server/internal/models"
	"server/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AdminRepo repository.AdminRepository
	UserRepo  repository.UserRepository
}

type CustomClaims struct {
	AdminID int `json:"admin_id"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte("VhakmGcXVPsQ05YyxKsj4SNdODjyK8kx")

func NewAuthService(adminRepo repository.AdminRepository, userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		AdminRepo: adminRepo,
		UserRepo:  userRepo,
	}
}

func (s *AuthService) AdminLogin(username, password string) (*models.Admin, error) {
	admin, err := s.AdminRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if admin.Password != password {
		return nil, config.ErrInvalidPassword
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

func (a *AuthService) GenerateToken(adminID int) (string, error) {
	claims := CustomClaims{
		AdminID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *AuthService) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
