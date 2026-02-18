package todo

import "Todo-list/domain"

func (svc *service) Create(todo domain.Todos) (*domain.Todos, error) {
	return svc.todoRepo.Create(todo)

}
