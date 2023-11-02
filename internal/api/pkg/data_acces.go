package pkg

import (
	"UserAPI/internal/api/utils"
	"UserAPI/internal/data"
	"UserAPI/internal/models"

	"log"
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

func ValidatheUser(users models.LoginUsers) (bool, error) {
	if ExistsEmail(users.Email) {
		log.Println("email already exists")
	}

	validacredential, err := db.Query("SELECT Password FROM Users WHERE Email = ?", users.Email)
	if err != nil {
		return false, err
	}
	defer validacredential.Close()
	var storePassword string
	found := false

	for validacredential.Next() {
		if err := validacredential.Scan(&storePassword); err != nil {
			return false, err
		}

		if utils.CheckP(users.Password, storePassword) {
			found = true
			break
		}

	}

	return found, nil
}

func DeleteUser(token []string) {
	decodetoken := ""
	qdeleteuser := "DELETE FROM Users WHERE Email = ?"
	_, err := db.Exec(qdeleteuser, decodetoken)
	if err != nil {
		log.Println("Error deleting")
	}

}
