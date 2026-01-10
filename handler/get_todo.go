package handler

import (
	"Todo-list/todo"
	"Todo-list/util"
	"net/http"
	"strconv"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	tId := r.PathValue("id")
	id, err := strconv.Atoi(tId)
	if err != nil {
		http.Error(w, "please give me valid id", 400)
		return
	}
	for _, t := range todo.Todolist {
		if t.ID == id {
			util.SendData(w, t, 200)

			// encode := json.NewEncoder(w)
			// encode.Encode(todo)
			return

		}
	}
}
