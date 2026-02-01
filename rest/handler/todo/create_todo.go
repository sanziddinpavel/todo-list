package todo

import (
	"Todo-list/repo"
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateTodo struct {
	Text        string `json:"text"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

func (h *Handler) CreateTodos(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateTodo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return

	}

	createdTodo, err := h.todoRepo.Create(repo.Todos{
		Text:        req.Text,
		Description: req.Description,
		IsDone:      req.IsDone,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, createdTodo, http.StatusCreated)

}
