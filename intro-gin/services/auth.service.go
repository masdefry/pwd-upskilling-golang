package services

import (
	"intro-gin/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateToken(Id string, Role models.UserRole, JwtExpireTime time.Duration, JwtSecret string)(string, error) {
	claims := jwt.MapClaims{
		"userId": Id,
		"role":    Role,
		"exp":     time.Now().Add(JwtExpireTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(JwtSecret))
}