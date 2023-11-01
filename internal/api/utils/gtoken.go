package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = []byte("secret-key")

func Tokenauth(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SECRET_KEY)
	return tokenString, err
}
