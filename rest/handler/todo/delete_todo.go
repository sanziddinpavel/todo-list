package todo

import (
	"Todo-list/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteTodos(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	tID, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid product id", 400)
		return
	}
	err = h.todoRepo.Delete(tID)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "successfully deleted", http.StatusOK)
}
