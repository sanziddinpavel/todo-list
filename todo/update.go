package todo

import "Todo-list/domain"

func (svc *service) Update(t domain.Todos) (*domain.Todos, error) {
	return svc.todoRepo.Update(t)

}
