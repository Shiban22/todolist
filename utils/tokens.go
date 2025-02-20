package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateToken generates a JWT token.
func GenerateToken(secretKey string, userID int64, expireDuration time.Duration) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expireDuration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   string(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken parses a JWT token.
func ParseToken(secretKey, tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
