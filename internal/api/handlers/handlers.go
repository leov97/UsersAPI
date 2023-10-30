package handlers

import (
	utils "UserAPI/internal/api/utils"
	"encoding/json"
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
}
