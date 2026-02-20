package todo

import (
	"Todo-list/domain"
	todoHndlr "Todo-list/rest/handler/todo"
)

type Service interface {
	todoHndlr.Service
}

type TodoRepo interface {
	Create(t domain.Todos) (*domain.Todos, error)
	Get(todoID int) (*domain.Todos, error)
	Count() (int64, error)
	List(page, limit int64) ([]*domain.Todos, error)
	Delete(todoID int) error
	Update(t domain.Todos) (*domain.Todos, error)
}
