package pkg

import (
	"UserAPI/internal/api/utils"
	"UserAPI/internal/models"

	"log"
)

var db = models.Datasql()

func PostNewUser(datos utils.NewUser) {
	if ExistsEmail(datos.Email) {
		log.Fatal("There is already an account with that email")
	}
	queryregister := "INSERT INTO Users (Username, Email, Password) VALUES (?, ?, ?)"
	_, err := db.Exec(queryregister, datos.Name, datos.Email, datos.Password)
	if err != nil {
		log.Fatal(err)
	}

}

func ExistsEmail(email string) bool {

	compare, err := db.Query("SELECT Email FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer compare.Close()

	useremails := make([]string, 0)
	for compare.Next() {
		var Email string
		if err := compare.Scan(&Email); err != nil {
			log.Fatal(err)
		}
		useremails = append(useremails, Email)
	}

	if err := compare.Err(); err != nil {
		log.Fatal(err)
	}

	for _, emails := range useremails {
		if email == emails {
			return true
		}
	}

	return false

}
