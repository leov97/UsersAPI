package auth

import "github.com/dgrijalva/jwt-go"

func ValidateJWT(tokenString string, secretKey string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
