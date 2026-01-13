package handler

import (
	"Todo-list/database"
	"Todo-list/util"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteTodos(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	tID, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid product id", 400)
		return
	}
	database.Delete(tID)

	util.SendData(w, "successfully deleted", http.StatusOK)
}
