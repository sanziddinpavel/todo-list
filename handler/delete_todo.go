package handler

import (
	"Todo-list/todo"
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

	for i, t := range todo.Todolist {
		if t.ID == tID {
			todo.Todolist = append(todo.Todolist[:i], todo.Todolist[i+1:]...)
			http.Error(w, "todo deleted", 200)
		}

	}

}
