package utils

import (
	"Shopping-cart/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id_user":  id,
		"username": "test",
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}
