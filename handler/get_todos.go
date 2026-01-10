package handler

import (
	"Todo-list/todo"
	"Todo-list/util"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}
	util.SendData(w, todo.Todolist, 200)

	// encoder := json.NewEncoder(w)
	// encoder.Encode(Todolist)

}
