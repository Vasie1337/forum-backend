package services

import (
	"errors"
	"server/internal/config"
	"server/internal/models"
	"server/internal/repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	AdminRepo repository.AdminRepository
	UserRepo  repository.UserRepository
}

type CustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

var (
	adminJWTSecretKey = []byte("thrdgf$jhnr6ju3$567h6rtj")
	userJWTSecretKey  = []byte("r56$jjr5y567t6fy$jrtyjyt")
)

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

	if user.Password != password {
		return nil, config.ErrInvalidPassword
	}

	return user, nil
}

func (a *AuthService) GenerateAdminToken(adminID int) (string, error) {
	claims := CustomClaims{
		ID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(adminJWTSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *AuthService) GenerateUserToken(userID int) (string, error) {
	claims := CustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(userJWTSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func cleanTokenString(token string) (string, error) {
	token = strings.TrimSpace(token)

	//for _, char := range token {
	//	if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '.' || char == '-') {
	//		return "", errors.New("invalid token: contains illegal characters")
	//	}
	//}

	return token, nil
}

func (a *AuthService) ValidateAdminToken(tokenString string) (*CustomClaims, error) {
	cleanedToken, err := cleanTokenString(tokenString)
	if err != nil {
		return nil, err
	}
	return a.validateToken(cleanedToken, adminJWTSecretKey)
}

func (a *AuthService) ValidateUserToken(tokenString string) (*CustomClaims, error) {
	cleanedToken, err := cleanTokenString(tokenString)
	if err != nil {
		return nil, err
	}
	return a.validateToken(cleanedToken, userJWTSecretKey)
}

func (a *AuthService) validateToken(tokenString string, secretKey []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
