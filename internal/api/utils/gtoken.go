package utils

import (
	"UserAPI/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
)

// var SECRET_KEY = []byte(env.SecreKey.Key)

func Tokenauth(id string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	env := config.NewDatabaseConfig()
	key := []byte(env.SecreKey.Key)
	claims.Valid()
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	tokenString, err := token.SignedString(key)
	return tokenString, err
}
