package todo

import (
	"Todo-list/repo"
	"Todo-list/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	todoRepo    repo.TodoRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	todoRepo repo.TodoRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		todoRepo:    todoRepo,
	}

}
