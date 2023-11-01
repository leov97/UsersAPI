package routes

import (
	"UserAPI/internal/api/handlers"
	"net/http"
)

func RegisterRoute() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.RegisterUserHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func LoginRoute() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.LoginUsersHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
