package handlers

import (
	"UserAPI/internal/api/pkg"
	utils "UserAPI/internal/api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func decode(r *http.Request, v interface{}) {
	data := r.Body
	decode := json.NewDecoder(data)

	decode.Decode(v)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &utils.NewUser{}
	decode(r, user)

	w.Write([]byte(user.Email))
	w.Write([]byte(user.Name))
	w.Write([]byte(user.Password))
	hash, _ := utils.Gepass(user.Password)

	user.Password = hash

	pkg.PostNewUser(*user)
}

func LoginUsersHandler(w http.ResponseWriter, r *http.Request) {
	loginuser := &utils.LoginUsers{}
	decode(r, loginuser)

	w.Write([]byte(loginuser.Email))
	w.Write([]byte(loginuser.Password))
	// hash, _ := utils.Gepass(loginuser.Password)
	// loginuser.Password = hash
	fmt.Println(loginuser.Email)

	userValid, err := pkg.ValidatheUser(*loginuser)
	log.Println(userValid)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !userValid {
		http.Error(w, "Email or Password is not valid", http.StatusUnauthorized)
		return
	}

	token, err := utils.Tokenauth(loginuser.Email)

	if err != nil {
		log.Println("token is not valid", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login suseful"))
	log.Println(token)
}
