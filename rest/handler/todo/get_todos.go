package todo

import (
	"Todo-list/database"
	"Todo-list/util"
	"net/http"
)

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}
	util.SendData(w, database.List(), 200)

	// encoder := json.NewEncoder(w)
	// encoder.Encode(Todolist)

}
