package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID string, jwtSecret string) (string, error) {
	claim := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ParseToken(accessToken string, jwtSecret string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	userID := token.Claims.(jwt.MapClaims)["sub"].(string)
	return userID, nil
}
