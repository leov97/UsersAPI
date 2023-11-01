package utils

import "golang.org/x/crypto/bcrypt"

func Gepass(gpass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(gpass), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}
