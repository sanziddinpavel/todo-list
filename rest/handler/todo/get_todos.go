package todo

import (
	"Todo-list/util"
	"net/http"
)

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todoList, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	util.SendData(w, todoList, http.StatusOK)

	// encoder := json.NewEncoder(w)
	// encoder.Encode(Todolist)

}
