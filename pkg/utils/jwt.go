package utils

import (
	"fmt"
	"time"

	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

func GenerateJWT(config *config.Config, username, password string) (string, error) {
	claims := Claims{username, password, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    "FileHub",
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.Credentials.Secret)

	return token.SignedString(secret)
}

func ValidateToken(config *config.Config, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Credentials.Secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("error parsing token claims")
	}

	// Check if the token is expired
	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("expiration time not found in claims")
	}
	expirationTime := time.Unix(int64(exp), 0)
	if time.Now().After(expirationTime) {
		return fmt.Errorf("token has expired")
	}

	return nil
}
