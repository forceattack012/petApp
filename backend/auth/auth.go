package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func AccessToken(signature []byte, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		Audience:  username,
	})

	ss, err := token.SignedString(signature)
	if err != nil {
		return "", err
	}
	return ss, nil
}
