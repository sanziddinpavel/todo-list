package todo

import (
	"Todo-list/repo"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateTodos(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // Go 1.22+ pattern: /todos/{id}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	var todo repo.Todos
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	todo.ID = id

	if _, err := h.todoRepo.Update(todo); err != nil {
		http.Error(w, "failed to update todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
