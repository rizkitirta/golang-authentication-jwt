package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "jwt-service",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "s3creet!!#!"
	}

	return secretKey
}

func (s *jwtService) GenerateToken(userId string) (string, error) {
	claims := &jwtCustomClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, err := t.Method.(*jwt.SigningMethodHMAC); err != nil {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
