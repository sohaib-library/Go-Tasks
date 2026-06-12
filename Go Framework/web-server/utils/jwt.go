package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {

	claims := jwt.MapClaims{
		
		"email":   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")

	

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}