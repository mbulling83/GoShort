package auth

import (
	"GoShort/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	jwtSecret := []byte(config.Get("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates the provided JWT and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	jwtSecret := []byte(config.Get("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
