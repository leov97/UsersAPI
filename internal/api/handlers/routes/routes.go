package routes

import (
	"UserAPI/internal/api/handlers"
	"net/http"
)

func RegisterRoute() {
	http.HandleFunc("/register", handlers.RegisterUserHandler)
}
