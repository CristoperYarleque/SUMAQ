package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Utils struct{}

type UtilsInterface interface {
	HashPassword(password string) (string, error)
	CheckPassword(hash string, password string) error
	GenerateJWT(userID int, role string) (string, error)
	ParseJWT(tokenStr string) (jwt.MapClaims, error)
}

var jwtKey = []byte("sumaq-2025")

// HashPassword genera un hash seguro
func (u *Utils) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compara una contraseña con su hash
func (u *Utils) CheckPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateJWT genera un token firmado
func (u *Utils) GenerateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseJWT valida y extrae los claims
func (u *Utils) ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}
