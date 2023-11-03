package main

import (
	"UserAPI/internal/api/routes"
	"log"
	"net/http"
)

func main() {

	//db  connection.Datasql()
	routes.RegisterRoute()
	routes.LoginRoute()
	routes.LogoutRoute()
	routes.DeleteRoute()
	starlistener()

}

func starlistener() {
	log.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
