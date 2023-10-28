package main

import (
	connection "UserAPI/internal/config"
	"fmt"
	"log"
)

func main() {
	db := connection.Datasql()
	rows, err := db.Query("SELECT ID, Nombre FROM Register")
	if err != nil {
		log.Fatal(err)
		// Maneja el error de manera apropiada para tu API
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var nombre string
		if err := rows.Scan(&id, &nombre); err != nil {
			log.Fatal(err)
			// Maneja el error de manera apropiada para tu API
		}
		fmt.Printf("ID: %d, Nombre: %s\n", id, nombre)
	}

}
