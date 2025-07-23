package utils

import (
	"gin-mvc/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", ""))

func GenerateJWT(userID uint, email string, name string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // token hết hạn sau 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
