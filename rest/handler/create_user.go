package handler

import (
	"Todo-list/database"
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var NewUsers database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewUsers)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return

	}

	createdUser := NewUsers.Store()

	util.SendData(w, createdUser, http.StatusCreated)

}
