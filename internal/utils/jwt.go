package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Konstanta untuk durasi token
const TokenDuration = time.Minute * 50

// Fungsi untuk menghasilkan token JWT
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(TokenDuration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Fungsi untuk mendekode token JWT dan mendapatkan user ID
func DecodeToken(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Periksa apakah token sudah kedaluwarsa
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return 0, errors.New("token has expired")
		}
	} else {
		return 0, errors.New("invalid token expiration")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}
	return uint(userID), nil
}
