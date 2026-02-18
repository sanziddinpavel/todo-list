package todo

import "Todo-list/domain"

type Service interface {
	Create(domain.Todos) (*domain.Todos, error)
	Get(id int) (*domain.Todos, error)
	List() ([]*domain.Todos, error)
	Update(domain.Todos) (*domain.Todos, error)
	Delete(id int) error
}
