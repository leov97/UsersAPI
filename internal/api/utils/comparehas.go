package utils

import "golang.org/x/crypto/bcrypt"

func CheckP(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
