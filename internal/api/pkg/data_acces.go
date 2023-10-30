package pkg

import (
	"UserAPI/internal/api/utils"
	"UserAPI/internal/models"

	"log"
)

var db = models.Datasql()

func PostNewUser(datos utils.NewUser) {
	queryregister := "INSERT INTO Register (Name, Email, Password) VALUES (?, ?, ?)"
	_, err := db.Exec(queryregister, datos.Name, datos.Email, datos.Password)
	if err != nil {
		log.Fatal(err)
	}

}

func ExistsEmail() bool {

	rows, err := db.Query("SELECT ID, Nombre FROM Register")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return true

}
