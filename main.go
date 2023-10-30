package main

import (
	"UserAPI/internal/api/handlers/routes"
	"log"
	"net/http"
)

func main() {
	//db := connection.Datasql()
	routes.RegisterRoute()
	starlistener()
}

func starlistener() {
	log.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
