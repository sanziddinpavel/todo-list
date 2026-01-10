package handler

import (
	"Todo-list/todo"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateTodos(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	tID, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid product id", 400)
		return
	}

	for _, todo := range todo.Todolist {
		if todo.ID == tID {

		}
		http.Error(w, "todo pai ni", 400)

	}

}
