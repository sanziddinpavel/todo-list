package todo

import (
	"Todo-list/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	tId := r.PathValue("id")
	id, err := strconv.Atoi(tId)
	if err != nil {
		http.Error(w, "please give me valid id", 400)
		return
	}
	product, err := h.todoRepo.Get(id)
	if err != nil {
		http.Error(w, "Give me a valid todo id", http.StatusInternalServerError)
	}
	if product == nil {
		util.SendError(w, 404, "Product not find")

	}

	util.SendData(w, product, 200)

	// encode := json.NewEncoder(w)
	// encode.Encode(todo)
	return

}
