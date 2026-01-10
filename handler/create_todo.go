package handler

import (
	"Todo-list/todo"
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTodos(w http.ResponseWriter, r *http.Request) {

	var NewTodos todo.Todos
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewTodos)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return

	}

	NewTodos.ID = len(todo.Todolist) + 1
	todo.Todolist = append(todo.Todolist, NewTodos)
	util.SendData(w, NewTodos, 200)
	// incoder := json.NewEncoder(w)
	// incoder.Encode(NewTodos)

}
