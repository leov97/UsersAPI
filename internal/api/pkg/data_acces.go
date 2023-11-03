package pkg

import (
	"UserAPI/internal/api/auth"
	"UserAPI/internal/api/utils"
	"UserAPI/internal/data"
	"UserAPI/internal/models"
	"strconv"

	"log"

	"github.com/dgrijalva/jwt-go"
)

var db = data.Datasql()

// sthis function send tha data to the database
func PostNewUser(datos models.NewUser) {

	if ExistsEmail(datos.Email) {
		log.Panic("There is already an account with that email")
	}
	queryregister := "INSERT INTO Users (Username, Email, Password) VALUES (?, ?, ?)"
	_, err := db.Exec(queryregister, datos.Name, datos.Email, datos.Password)
	if err != nil {
		log.Panic(err)
	}

}

// this function verifies that the email is valid
func ExistsEmail(email string) bool {

	compare, err := db.Query("SELECT Email FROM Users")
	if err != nil {
		log.Panic(err)
	}
	defer compare.Close()

	useremails := make([]string, 0)
	for compare.Next() {
		var Email string
		if err := compare.Scan(&Email); err != nil {
			log.Panic(err)
		}
		useremails = append(useremails, Email)
	}

	if err := compare.Err(); err != nil {
		log.Panic(err)

	}

	for _, emails := range useremails {
		if email == emails {
			return true
		}
	}

	return false
}

func ValidatheUser(users models.LoginUsers) (string, bool, error) {
	if ExistsEmail(users.Email) {
		log.Println("email already exists")
	}

	validacredential, err := db.Query("SELECT UserID, Password FROM Users WHERE Email = ?", users.Email)
	if err != nil {
		return "", false, err
	}
	defer validacredential.Close()
	var storePassword string
	found := false
	var userID = models.Authlogin{}

	for validacredential.Next() {
		if err := validacredential.Scan(&userID.ID, &storePassword); err != nil {
			return "", false, err
		}

		if utils.CheckP(users.Password, storePassword) {
			found = true
			break
		}
	}

	return userID.ID, found, nil
}

func CheckLogin(tokenStr string, secretkey string) (int, string, error) {
	claims := &models.Authlogin{}

	_, err := auth.ValidateJWT(tokenStr, secretkey)
	if err != nil {
		log.Println("Error validating", err)
		return 0, "", err
	}

	jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println("Error al convertir el string a int:", err)

	}

	return id, claims.User, nil
}

func PostUserActives(id int, email string) {

	queryregister := "INSERT INTO UsersActive (Id, Email) VALUES (?, ?)"
	_, err := db.Exec(queryregister, id, email)
	if err != nil {
		log.Println(err)
	}

}

func DeleteUserActive(email string) {
	queryDelete := "DELETE FROM UsersActive WHERE Email = ?"
	_, err := db.Exec(queryDelete, email)
	if err != nil {
		log.Println(err)
	}
}

func UserEnabled(email string) (int, error) {

	validacredential, err := db.Query("SELECT Id FROM UsersActive WHERE Email = ?", email)
	if err != nil {
		return 0, err
	}
	defer validacredential.Close()

	var id int

	for validacredential.Next() {
		if err := validacredential.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func DeleteUser(id int) {
	queryDelete := "DELETE FROM Users WHERE UserID = ?"
	_, err := db.Exec(queryDelete, id)
	if err != nil {
		log.Println(err)
	}
}
