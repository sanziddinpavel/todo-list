package handler

import (
	"Todo-list/config"
	"Todo-list/database"
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return

	}

	user := database.Find(reqLogin.Email, reqLogin.Password)

	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		http.Error(w, "failed to create jwt", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)

}
