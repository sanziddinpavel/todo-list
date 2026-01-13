package handler

import (
	"Todo-list/database"
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTodos(w http.ResponseWriter, r *http.Request) {

	var NewTodos database.Todos
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewTodos)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return

	}

	createdTodo := database.Store(NewTodos)

	util.SendData(w, createdTodo, 200)
	// incoder := json.NewEncoder(w)
	// incoder.Encode(NewTodos)

}
