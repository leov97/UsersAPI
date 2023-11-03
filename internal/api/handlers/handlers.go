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

// var activeTokens = make(map[string]bool)
var env = config.NewDatabaseConfig()
var key = []byte(env.SecreKey.Key)

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
	ActiveUsers := &models.ActiveUsers{

		Users: make(map[string]string),
	}
	tokenuseractive := &models.Authlogin{}
	decode(r, loginuser)

	w.Write([]byte(loginuser.Email))
	w.Write([]byte(loginuser.Password))
	// hash, _ := utils.Gepass(loginuser.Password)
	// loginuser.Password = hash
	fmt.Println(loginuser.Email)

	userID, userValid, err := pkg.ValidatheUser(*loginuser)
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

	token, err := utils.Tokenauth(userID, loginuser.Email)
	tokenuseractive.Token = token
	if err != nil {
		log.Println("Error al verificar el usuario:", err)
		return
	}

	id, user, err := pkg.CheckLogin(token, string(key))

	pkg.PostUserActives(id, user)

	if err != nil {
		log.Println("token is not valid", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login suseful"))
	log.Println(token)
	log.Println(ActiveUsers.Users)
}

// this function logout the user from the session
func LogoutUser(w http.ResponseWriter, r *http.Request) {

	tokenuser := &models.Authlogin{}
	decode(r, tokenuser)
	env := config.NewDatabaseConfig()
	key := []byte(env.SecreKey.Key)

	_, err := auth.ValidateJWT(tokenuser.Token, string(key))
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not Token"))
	} else {
		_, email, err := pkg.CheckLogin(tokenuser.Token, string(key))
		if err != nil {
			log.Println(err)
		}

		pkg.DeleteUserActive(email)
		log.Println("Logout Suseful")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logout Suseful"))
	}

}

func DeleteAccouunt(w http.ResponseWriter, r *http.Request) {
	deleteUser := &models.Authlogin{}
	decode(r, deleteUser)
	env := config.NewDatabaseConfig()
	key := []byte(env.SecreKey.Key)

	_, err := auth.ValidateJWT(deleteUser.Token, string(key))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Token"))
		return
	}
	_, email, err := pkg.CheckLogin(deleteUser.Token, string(key))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error getting user email"))
		return
	}

	IdActive, err := pkg.UserEnabled(email)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not find active user"))
		return
	}

	pkg.DeleteUser(IdActive)

	pkg.DeleteUserActive(email)

	log.Println("Successfully deleted user")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted user"))
}
