package todo

import "Todo-list/domain"

type Service interface {
	Create(domain.Todos) (*domain.Todos, error)
	Get(id int) (*domain.Todos, error)
	List(page, limit int64) ([]*domain.Todos, error)
	Count() (int64, error)
	Update(domain.Todos) (*domain.Todos, error)
	Delete(id int) error
}
