package todo

import "Todo-list/domain"

func (svc *service) List() ([]*domain.Todos, error) {
	return svc.todoRepo.List()
}
