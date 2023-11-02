package handlers

import (
	"UserAPI/internal/api/auth"
	"UserAPI/internal/api/pkg"
	utils "UserAPI/internal/api/utils"
	"UserAPI/internal/config"
	"UserAPI/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var activeTokens = make(map[string]bool)

func decode(r *http.Request, v interface{}) {
	data := r.Body
	decode := json.NewDecoder(data)

	decode.Decode(v)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.NewUser{}
	decode(r, user)

	w.Write([]byte(user.Email))
	w.Write([]byte(user.Name))
	w.Write([]byte(user.Password))
	hash, _ := utils.Gepass(user.Password)

	user.Password = hash

	pkg.PostNewUser(*user)
}

func LoginUsersHandler(w http.ResponseWriter, r *http.Request) {
	loginuser := &models.LoginUsers{}
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
	activeTokens[token] = true

	log.Println(activeTokens)
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

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	token := &models.Authlogin{}
	decode(r, token)

	env := config.NewDatabaseConfig()
	key := []byte(env.SecreKey.Key)
	_, err := auth.ValidateJWT(token.Token, string(key))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not Token"))
	} else {
		log.Println("Logout Suseful")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logout Suseful"))
	}

}

func DeleteAccouunt(w http.ResponseWriter, r *http.Request) {
	tokenuser := &models.Authlogin{}

	if tokenuser.Token != "" {

		delete(activeTokens, tokenuser.Token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not Token"))
	}
}
