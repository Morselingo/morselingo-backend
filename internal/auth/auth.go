package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey []byte
var tokenExpiration = time.Hour * 24

type ContextKey string

const (
	ContextKeyUser ContextKey = "authenticatedUser"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func InitializeAuthentication(secretKey string) error {
	if secretKey == "" {
		return fmt.Errorf("JWT secret key cannot be empty")
	}
	jwtSecretKey = []byte(secretKey)
	return nil
}

func GenerateToken(username string) (string, error) {
	currentTime := time.Now()
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
			NotBefore: jwt.NewNumericDate(currentTime),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}
